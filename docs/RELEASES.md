# Releases - Sürüm Yönetimi

Bu dokümanda release süreci, sürüm numaralandırması ve CLI komutları açıklanmaktadır.

## 📦 Mevcut Sürümler

Tüm sürümler [GitHub Releases](https://github.com/ozanbozkurtt/perf-analyzer/releases) sayfasında mevcuttur.

Her release için aşağıdaki dosyalar bulunur:

```
v1.0.0/
├── benchmark-linux-arm64           # Linux ARM64 (Raspberry Pi, etc.)
├── benchmark-linux-x86_64          # Linux x86_64 (Normal bilgisayar)
├── benchmark-macos-arm64           # macOS (Apple Silicon M1/M2/M3)
├── benchmark-macos-x86_64          # macOS (Intel)
├── benchmark-windows-x86_64.exe    # Windows
└── checksums.txt                   # SHA256 kontrol toplamları
```

## 🚀 Release Süreci

### Otomatik Release (GitHub Actions)

1. **Tag oluşturun:**
```bash
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

2. **Otomatik işlemler başlayır:**
   - GitHub Actions workflow tetiklenir
   - Tüm platformlar için binary'ler derlenecek
   - Checksum'lar oluşturulacak
   - Release sayfasında yayınlanacak

3. **Release sayfası güncellenir:**
   - Binary'ler yüklenir
   - Checksum'lar eklenir
   - Release notes gösterilir

### Manuel Release (GoReleaser)

GoReleaser kuruysa:

```bash
# Draft mod (test etmek için)
goreleaser release --skip=publish --skip=validate

# Production release
goreleaser release
```

## 📝 Sürüm Numaralandırması

[Semantic Versioning](https://semver.org/) standardını takip ederiz:

```
v<MAJOR>.<MINOR>.<PATCH>
```

### Örnekler

- **v1.0.0** - İlk release (major=1, minor=0, patch=0)
- **v1.1.0** - Yeni özellik eklendi (minor artırıldı)
- **v1.0.1** - Bug fix (patch artırıldı)
- **v2.0.0** - Breaking changes (major artırıldı)

## 🏷️ Tag Oluşturmak

```bash
# Basit tag (annotated olmayan)
git tag v1.0.0

# Açıklamalı tag (önerilen)
git tag -a v1.0.0 -m "Release version 1.0.0"

# Açık mesaj ile
git tag -a v1.0.0 -m "Release version 1.0.0 - Add multi-platform support"

# Tag'ı push edin
git push origin v1.0.0

# Tüm tag'ları push edin
git push origin --tags
```

## ✅ Pre-Release Kontrol Listesi

Release yapmadan önce:

- [ ] Commit history'yi kontrol edin (`git log`)
- [ ] Tüm testleri çalıştırın (`go test ./...`)
- [ ] README.md'yi güncelleyin
- [ ] Changelog'u hazırlayın
- [ ] Sürüm numarasını belirleyin
- [ ] go.mod versiyonunu kontrol edin
- [ ] Tag'ı push edin

## 🔍 Release Doğrulaması

### Checksum Doğrulama

```bash
# checksums.txt dosyasını indirin
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/v1.0.0/checksums.txt

# Tüm dosyaları doğrulayın
sha256sum -c checksums.txt

# Tek dosya doğrulama
echo "abc123... benchmark-macos-arm64" | sha256sum -c
```

### Binary Test Etmek

```bash
# İndirilen binary'yi çalıştırın
./benchmark-macos-arm64 --version

# Benchmark'i çalıştırın
./benchmark-macos-arm64
```

## 📋 Changelog Formatı

Commit mesajları otomatik olarak changelog'a dönüştürülür.

### Commit Konvensiyonları

```bash
# Feature (yeni özellik)
git commit -m "feat: add new performance metric"

# Fix (hata düzeltme)
git commit -m "fix: handle edge case in memory calculation"

# Docs (dokümantasyon)
git commit -m "docs: update installation guide"

# Chore (bakım işleri, changelog'a girmez)
git commit -m "chore: update dependencies"

# Perf (performans iyileştirmesi)
git commit -m "perf: optimize cpu benchmark"
```

## 🔄 Eski Sürümleri Yönetmek

### Sürümleri Listele

```bash
# Local
git tag

# Remote
git ls-remote --tags origin
```

### Sürüm Ayrıntıları

```bash
git show v1.0.0
```

### Sürüm Arasında Fark

```bash
git diff v1.0.0..v1.1.0
```

## 🗑️ Sürümü Silmek

**Dikkat:** GitHub'da yayınlandıysa silemezsiniz (best practices).

### Local Tag Silme

```bash
git tag -d v1.0.0
```

### Remote Tag Silme (Yayınlananlar silinmez)

```bash
git push origin :refs/tags/v1.0.0
```

## 🌐 Release Sayfasında Ne Gösterilir

GitHub Actions otomatik olarak aşağıdaları oluşturur:

1. **Binary'ler**
   - 5 platform için executable dosyalar
   - Boyut ve hash'le gösterilir

2. **Checksum Dosyası**
   - SHA256 kontrol toplamları
   - Güvenlik doğrulaması için

3. **Release Notes**
   - Commit history'den otomatik oluşturulur
   - Kategorilere ayrılmış (feat, fix, docs, etc.)

4. **Install Komutları**
   - Hızlı kurulum scriptleri
   - Platform bazlı instrüksiyonlar

## 📊 Release İstatistikleri

```bash
# Total releases
git tag | wc -l

# Total commits
git rev-list --count HEAD

# Commits since last release
git log v1.0.0..HEAD --oneline | wc -l
```

## 🔐 GPG İmzalama (İsteğe Bağlı)

Ek güvenlik için tag'ları GPG ile imzalayabilirsiniz:

```bash
# İmzalı tag oluştur
git tag -s v1.0.0 -m "Release version 1.0.0"

# Imzayı kontrol et
git tag -v v1.0.0
```

## 📞 Yardım

- [GitHub Releases](https://github.com/ozanbozkurtt/perf-analyzer/releases) - Tüm sürümler
- [GitHub Actions](https://github.com/ozanbozkurtt/perf-analyzer/actions) - Build logs
- [GoReleaser Docs](https://goreleaser.com/) - Release automation

---

**Sorunuz varsa GitHub Issues'ı kullanın!** 🚀
