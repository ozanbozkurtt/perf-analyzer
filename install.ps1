# Perf Analyzer Installer for Windows PowerShell

param(
    [string]$InstallDir = "$env:ProgramFiles\PerfAnalyzer"
)

$ErrorActionPreference = "Stop"

# Renklar
$Colors = @{
    Blue   = [System.ConsoleColor]::Cyan
    Green  = [System.ConsoleColor]::Green
    Yellow = [System.ConsoleColor]::Yellow
    Red    = [System.ConsoleColor]::Red
}

function Write-Color([string]$Message, [System.ConsoleColor]$Color = $Colors.Blue) {
    Write-Host $Message -ForegroundColor $Color
}

function Write-Success([string]$Message) {
    Write-Color "✓ $Message" $Colors.Green
}

function Write-Error-Custom([string]$Message) {
    Write-Color "✗ $Message" $Colors.Red
}

function Write-Warning-Custom([string]$Message) {
    Write-Color "⚠️  $Message" $Colors.Yellow
}

function Write-Info([string]$Message) {
    Write-Color "ℹ️  $Message" $Colors.Blue
}

# Başlık
Write-Host ""
Write-Color "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" $Colors.Blue
Write-Color "  Perf Analyzer Installer (Windows)" $Colors.Blue
Write-Color "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" $Colors.Blue
Write-Host ""

# Sabitler
$GITHUB_REPO = "ozanbozkurtt/perf-analyzer"
$GITHUB_API = "https://api.github.com/repos/$GITHUB_REPO"
$BINARY_NAME = "benchmark.exe"

# En son release'i bul
Write-Info "Fetching latest release..."
try {
    $response = Invoke-WebRequest -Uri "$GITHUB_API/releases/latest" -UseBasicParsing
    $release = $response.Content | ConvertFrom-Json
    $LATEST_RELEASE = $release.tag_name
} catch {
    Write-Error-Custom "Could not fetch latest release"
    exit 1
}

if (-not $LATEST_RELEASE) {
    Write-Error-Custom "Could not determine latest version"
    exit 1
}

Write-Success "Latest version: $LATEST_RELEASE"
Write-Host ""

# Kurulum dizinini oluştur
if (-not (Test-Path $InstallDir)) {
    Write-Info "Creating installation directory: $InstallDir"
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
}

# Binary'yi indir
Write-Info "Downloading binary..."
$BINARY_URL = "https://github.com/$GITHUB_REPO/releases/download/$LATEST_RELEASE/$BINARY_NAME"
$OUTPUT_PATH = Join-Path $InstallDir $BINARY_NAME

try {
    Invoke-WebRequest -Uri $BINARY_URL -OutFile $OUTPUT_PATH -UseBasicParsing
    Write-Success "Downloaded to $OUTPUT_PATH"
} catch {
    Write-Error-Custom "Download failed: $_"
    exit 1
}

# Checksum doğrula
Write-Info "Verifying checksum..."
try {
    $CHECKSUMS_URL = "https://raw.githubusercontent.com/$GITHUB_REPO/$LATEST_RELEASE/checksums.txt"
    $checksums = Invoke-WebRequest -Uri $CHECKSUMS_URL -UseBasicParsing | Select-Object -ExpandProperty Content

    # Checksum'ı hesapla
    $actualHash = (Get-FileHash -Path $OUTPUT_PATH -Algorithm SHA256).Hash.ToLower()

    # checksums.txt'de benchmark-windows-x86_64.exe arıyor
    $expectedHash = ($checksums | Select-String "benchmark-windows-x86_64" | Select-Object -First 1 | ForEach-Object { $_.Line.Split()[0] }).ToLower()

    if ($actualHash -eq $expectedHash) {
        Write-Success "Checksum verified"
    } else {
        Write-Warning-Custom "Could not verify checksum, continuing anyway..."
    }
} catch {
    Write-Warning-Custom "Checksum verification failed, continuing..."
}

Write-Host ""

# PATH'e ekle
Write-Info "Adding to PATH..."
$PATH = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($PATH -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$PATH;$InstallDir", "User")
    Write-Success "Added $InstallDir to PATH"
    Write-Warning-Custom "Please restart your PowerShell session for PATH changes to take effect"
} else {
    Write-Success "$InstallDir is already in PATH"
}

# Kurulum tamamlandı
Write-Host ""
Write-Color "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" $Colors.Blue
Write-Success "Perf Analyzer $LATEST_RELEASE is ready!"
Write-Color "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" $Colors.Blue
Write-Host ""

Write-Info "Usage:"
Write-Host "  benchmark              # Run benchmark" -ForegroundColor $Colors.Yellow
Write-Host ""
Write-Info "Location:"
Write-Host "  $OUTPUT_PATH" -ForegroundColor $Colors.Yellow
Write-Host ""
