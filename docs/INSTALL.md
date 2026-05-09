# Perf Analyzer - Kurulum Rehberi

Bu dokümanda Perf Analyzer'ı farklı işletim sistemlerine kurmak için adım adım talimatlar bulunmaktadır.

## 🚀 Hızlı Kurulum (Önerilen)

### macOS ve Linux

Tek komutla kurulum:

```bash
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash
```

**Ne yapılır:**
- En son sürümü GitHub Releases'ten indirir
- SHA256 checksum'ını doğrular
- `/usr/local/bin` dizinine kurar
- Gerekirse sudo izni ister

### Windows (PowerShell)

PowerShell'i Administrator olarak açıp çalıştırın:

```powershell
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/ozanbozkurtt/perf-analyzer/main/install.ps1'))
```

**Ne yapılır:**
- En son sürümü indirir
- `Program Files\PerfAnalyzer` dizinine kurar
- PATH'e otomatik olarak ekler

## 📥 Manuel Kurulum

### 1. Binary'yi İndirin

[GitHub Releases](https://github.com/ozanbozkurtt/perf-analyzer/releases) sayfasından cihazınıza uygun binary'yi seçin.

#### Platform Seçim

| Cihazınız | İndirilecek Dosya |
|-----------|------------------|
| macOS (Apple Silicon M1/M2/M3) | `benchmark-macos-arm64` |
| macOS (Intel) | `benchmark-macos-x86_64` |
| Linux (ARM64, Raspberry Pi) | `benchmark-linux-arm64` |
| Linux (x86_64, Normal bilgisayar) | `benchmark-linux-x86_64` |
| Windows | `benchmark-windows-x86_64.exe` |

### 2. macOS/Linux Kurulumu

```bash
# İndirilen dosyanın adını değiştirin (örnek)
cd ~/Downloads
chmod +x benchmark-macos-arm64

# /usr/local/bin'e kopyalayın (admin izni gerekebilir)
sudo cp benchmark-macos-arm64 /usr/local/bin/benchmark

# Kurulumu kontrol edin
benchmark --version
```

### 3. Windows Kurulumu

#### Seçenek A: Otomatik (PowerShell)

Yukarıdaki PowerShell install.ps1 scriptini kullanın.

#### Seçenek B: Manuel

1. `benchmark-windows-x86_64.exe` dosyasını indirin
2. `C:\Program Files\PerfAnalyzer` klasörü oluşturun
3. Dosyayı bu klasöre kopyalayın
4. PATH'e ekleyin:
   - Windows Start → "Environment Variables" yazın
   - "Edit the system environment variables" tıklayın
   - "Environment Variables..." düğmesini tıklayın
   - Path'i seçin → "Edit" tıklayın
   - "New" ile `C:\Program Files\PerfAnalyzer` ekleyin
   - OK'ye tıklayın ve PowerShell'i yeniden başlatın

## 🔍 Kurulumu Doğrulayın

```bash
# Sürümü kontrol edin
benchmark --version

# Benchmark'i çalıştırın
benchmark

# Yardım al
benchmark --help
```

## ✅ Kurulum Kontrol Listesi

- [ ] Binary'yi indirdim
- [ ] Dosyaya execute (çalıştırma) izni verdim (macOS/Linux)
- [ ] `/usr/local/bin` veya PATH'deki bir klasöre kopyaladım
- [ ] Terminal/PowerShell'i yeniden başlattım
- [ ] `benchmark --version` komutu çalışıyor
- [ ] `benchmark` komutu performans raporunu gösteriyor

## 🔧 Sorun Giderme

### "command not found: benchmark" hatası

**Sebep:** Binary PATH'de değil

**Çözüm:**
```bash
# Tüm PATH dizinlerini kontrol edin
echo $PATH

# Binary'nin konumunu bulun
which benchmark
# veya
find / -name benchmark 2>/dev/null

# Eğer bulunamazsa, dosyayı PATH'deki bir klasöre kopyalayın
cp ~/Downloads/benchmark-macos-arm64 /usr/local/bin/benchmark
```

### "Permission denied" hatası

**Sebep:** Dosya execute (çalıştırma) izni yok

**Çözüm:**
```bash
chmod +x /path/to/benchmark
```

### "quarantine" hatası (macOS)

**Sebep:** macOS internet'ten indirilen dosyayı kısıtlıyor

**Çözüm:**
```bash
# Karantina niteliğini kaldırın
xattr -d com.apple.quarantine /path/to/benchmark

# Veya Finder'da:
# 1. benchmark dosyasına sağ tıklayın
# 2. "Open" seçin
# 3. "Open" düğmesini tıklayın
```

### Windows'ta PowerShell hataları

**Hata:** "cannot be loaded because running scripts is disabled"

**Çözüm:**
```powershell
# PowerShell'i Administrator olarak açın
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser

# Sonra install scriptini çalıştırın
```

## 🔐 Checksum Doğrulaması (İsteğe Bağlı)

Güvenlik için indirilen dosyayı doğrulayabilirsiniz:

```bash
# checksums.txt dosyasını indirin
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/checksums.txt

# Doğrulama yapın
sha256sum -c checksums.txt

# Çıktı:
# benchmark-macos-arm64-v1.0.0: OK
```

## 🆚 Kaynak Koddan Derleme

Go 1.21 veya daha yenisini kurulu varsa:

```bash
# Repo'yu klonlayın
git clone https://github.com/ozanbozkurtt/perf-analyzer.git
cd perf-analyzer

# Derleyin
go build -o benchmark cmd/benchmark/main.go

# Kurulum
cp benchmark /usr/local/bin/
```

## 🔄 Güncelleme

Yeni sürüme güncellemek için:

```bash
# Kurulum scriptini tekrar çalıştırın
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash

# Sürümü kontrol edin
benchmark --version
```

## 🗑️ Kaldırma

### macOS/Linux

```bash
rm /usr/local/bin/benchmark
```

### Windows (PowerShell Admin)

```powershell
Remove-Item "C:\Program Files\PerfAnalyzer\benchmark.exe"
Remove-Item -Recurse "C:\Program Files\PerfAnalyzer"
```

## 📞 Destek

Kurulum sırasında sorun yaşıyorsanız:

1. Bu rehberi tekrar okuyun
2. [GitHub Issues](https://github.com/ozanbozkurtt/perf-analyzer/issues) açın
3. ozan.bozkurt@ode.al adresine e-posta gönderin

---

**Kurulum için başka sorularınız varsa lütfen iletişime geçin!** 🚀
