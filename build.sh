#!/bin/bash
set -euo pipefail

# Renkli çıktı
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PROJECT_NAME="perf-analyzer"
VERSION="${1:-dev}"
BINARY_NAME="benchmark"
OUTPUT_DIR="./dist"

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  ${PROJECT_NAME} Multi-Platform Build${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# Build hedefleri
declare -a PLATFORMS=(
  "linux:amd64:linux-x86_64"
  "linux:arm64:linux-arm64"
  "darwin:amd64:macos-x86_64"
  "darwin:arm64:macos-arm64"
  "windows:amd64:windows-x86_64"
)

# Çıkış dizinini oluştur
mkdir -p "${OUTPUT_DIR}"

# Derle
echo -e "${YELLOW}📦 Compiling binaries...${NC}\n"

for platform in "${PLATFORMS[@]}"; do
  IFS=':' read -r GOOS GOARCH NAME <<< "$platform"

  # Dosya ismi oluştur
  if [ "$GOOS" = "windows" ]; then
    OUTPUT_FILE="${OUTPUT_DIR}/${BINARY_NAME}-${NAME}-${VERSION}.exe"
  else
    OUTPUT_FILE="${OUTPUT_DIR}/${BINARY_NAME}-${NAME}-${VERSION}"
  fi

  echo -n "Building ${NAME}... "

  # Derle
  GOOS="$GOOS" GOARCH="$GOARCH" CGO_ENABLED=0 go build \
    -ldflags="-s -w -X main.Version=${VERSION} -X main.BuildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') -X main.Commit=$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown')" \
    -o "$OUTPUT_FILE" \
    ./cmd/benchmark

  # Dosya boyutunu al
  SIZE=$(du -h "$OUTPUT_FILE" | cut -f1)
  echo -e "${GREEN}✓ (${SIZE})${NC}"
done

echo -e "\n${YELLOW}📝 Creating checksums...${NC}"
cd "${OUTPUT_DIR}"
sha256sum benchmark-* > checksums.txt
cd ..

echo -e "\n${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}✓ Build completed successfully!${NC}"
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "\n${BLUE}📂 Output directory: ${OUTPUT_DIR}/${NC}"
ls -lh "${OUTPUT_DIR}"

echo -e "\n${BLUE}ℹ️  To install:${NC}"
echo -e "  ${YELLOW}macOS (arm64):${NC}   sudo cp dist/benchmark-macos-arm64-${VERSION} /usr/local/bin/benchmark"
echo -e "  ${YELLOW}macOS (x86_64):${NC}  sudo cp dist/benchmark-macos-x86_64-${VERSION} /usr/local/bin/benchmark"
echo -e "  ${YELLOW}Linux (arm64):${NC}   sudo cp dist/benchmark-linux-arm64-${VERSION} /usr/local/bin/benchmark"
echo -e "  ${YELLOW}Linux (x86_64):${NC}  sudo cp dist/benchmark-linux-x86_64-${VERSION} /usr/local/bin/benchmark"
echo -e "  ${YELLOW}Windows:${NC}        Copy dist/benchmark-windows-x86_64-${VERSION}.exe to a folder in PATH"
