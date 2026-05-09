#!/bin/bash
set -euo pipefail

# Renkli çıktı
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Sabitler
GITHUB_REPO="ozanbozkurtt/perf-analyzer"
GITHUB_API="https://api.github.com/repos/${GITHUB_REPO}"
INSTALL_DIR="${1:-/usr/local/bin}"
BINARY_NAME="benchmark"

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  Perf Analyzer Installer${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}\n"

# OS ve ARCH tespiti
OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
  Darwin)
    OS_NAME="macos"
    if [ "$ARCH" = "arm64" ]; then
      ARCH_NAME="arm64"
    elif [ "$ARCH" = "x86_64" ]; then
      ARCH_NAME="x86_64"
    else
      echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
      exit 1
    fi
    ;;
  Linux)
    OS_NAME="linux"
    if [ "$ARCH" = "aarch64" ]; then
      ARCH_NAME="arm64"
    elif [ "$ARCH" = "x86_64" ]; then
      ARCH_NAME="x86_64"
    else
      echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
      exit 1
    fi
    ;;
  *)
    echo -e "${RED}✗ Unsupported OS: $OS${NC}"
    exit 1
    ;;
esac

echo -e "${YELLOW}Detected system: ${BLUE}$OS ($ARCH)${NC}\n"

# En son release'i bul
echo -e "${YELLOW}📦 Fetching latest release...${NC}"
LATEST_RELEASE=$(curl -s "${GITHUB_API}/releases/latest" | grep tag_name | cut -d'"' -f4)

if [ -z "$LATEST_RELEASE" ]; then
  echo -e "${RED}✗ Could not fetch latest release${NC}"
  exit 1
fi

echo -e "${GREEN}✓ Latest version: ${LATEST_RELEASE}${NC}\n"

# Binary'yi indir
BINARY_URL="https://github.com/${GITHUB_REPO}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-${OS_NAME}-${ARCH_NAME}"

echo -e "${YELLOW}⬇️  Downloading binary...${NC}"
TEMP_FILE=$(mktemp)

if ! curl -sL "${BINARY_URL}" -o "${TEMP_FILE}"; then
  echo -e "${RED}✗ Download failed${NC}"
  rm -f "${TEMP_FILE}"
  exit 1
fi

# Dosya boyutu kontrol et
if [ ! -s "${TEMP_FILE}" ]; then
  echo -e "${RED}✗ Downloaded file is empty${NC}"
  rm -f "${TEMP_FILE}"
  exit 1
fi

# Checksum doğrula
echo -e "${YELLOW}🔐 Verifying checksum...${NC}"
CHECKSUMS=$(curl -s "https://github.com/${GITHUB_REPO}/releases/download/${LATEST_RELEASE}/checksums.txt")
EXPECTED_CHECKSUM=$(echo "$CHECKSUMS" | grep "benchmark-${OS_NAME}-${ARCH_NAME}" | awk '{print $1}')

if [ -z "$EXPECTED_CHECKSUM" ]; then
  echo -e "${YELLOW}⚠️  Checksum file not found, skipping verification${NC}"
else
  ACTUAL_CHECKSUM=$(shasum -a 256 "${TEMP_FILE}" | awk '{print $1}')
  if [ "$EXPECTED_CHECKSUM" != "$ACTUAL_CHECKSUM" ]; then
    echo -e "${RED}✗ Checksum mismatch!${NC}"
    echo -e "Expected: ${EXPECTED_CHECKSUM}"
    echo -e "Actual:   ${ACTUAL_CHECKSUM}"
    rm -f "${TEMP_FILE}"
    exit 1
  fi
  echo -e "${GREEN}✓ Checksum verified${NC}\n"
fi

# Binary'yi yükle
echo -e "${YELLOW}📝 Installing to ${INSTALL_DIR}/${BINARY_NAME}...${NC}"

# Yazma izni kontrol et
if [ ! -w "$INSTALL_DIR" ]; then
  echo -e "${YELLOW}⚠️  Need elevated privileges${NC}"
  sudo mv "${TEMP_FILE}" "${INSTALL_DIR}/${BINARY_NAME}"
  sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
else
  mv "${TEMP_FILE}" "${INSTALL_DIR}/${BINARY_NAME}"
  chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
fi

# Kurulum başarılı
echo -e "${GREEN}✓ Installation successful!${NC}\n"

# Sürüm kontrol
VERSION=$("${INSTALL_DIR}/${BINARY_NAME}" --version 2>/dev/null || echo "unknown")

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}✓ Perf Analyzer ${LATEST_RELEASE} is ready!${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}\n"

echo -e "${BLUE}Usage:${NC}"
echo -e "  ${YELLOW}${BINARY_NAME}${NC}              # Run benchmark"
echo -e "\n${BLUE}Location:${NC}"
echo -e "  ${YELLOW}${INSTALL_DIR}/${BINARY_NAME}${NC}\n"

# PATH kontrol et
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo -e "${YELLOW}⚠️  $INSTALL_DIR is not in your PATH${NC}"
  echo -e "   Add it with: ${BLUE}export PATH=\$PATH:${INSTALL_DIR}${NC}\n"
fi
