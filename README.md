# Performans Analiz Aracı (Perf Analyzer)

🚀 Donanım performansını derinlemesine analiz eden, kapsamlı raporlar üreten Go tabanlı bir benchmark ve sistem analiz aracı.

## Özellikler

- 🖥️ **CPU Analizi**: İşlemci performansı, çekirdek sayısı, frekans bilgileri
- 💾 **RAM Analizi**: Bellek kullanımı, sistem kapasitesi, alan tahmini
- 💿 **Disk Analizi**: Disk hızı, I/O performansı, kalan alan
- 🌐 **Ağ Analizi**: Bağlantı hızı ve durumu
- 📊 **Sistem Özeti**: Toplam sistem performans değerlendirmesi
- 🎯 **Kalite Puanlandırması**: Yapılandırılmış performans puanları
- 📈 **Öneriler**: Sistem iyileştirmesi için yapılandırılmış tavsiyeler
- 🎨 **Renkli Çıktı**: Terminal üzerinde kolay okunabilir formatlama

## Kurulum

### 🚀 Hızlı Kurulum (Derlenmiş Binary'ler)

Go yüklü olmadan direkt çalışabilen derlenmiş binary'leri kullanabilirsiniz.

#### macOS
```bash
# Apple Silicon (M1/M2/M3)
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash

# Intel
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash
```

#### Linux
```bash
# x86_64
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash

# ARM64
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash
```

#### Windows (PowerShell)
```powershell
# Administrator olarak PowerShell'i açın
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/ozanbozkurtt/perf-analyzer/main/install.ps1'))
```

#### Manuel Kurulum
[GitHub Releases](https://github.com/ozanbozkurtt/perf-analyzer/releases) sayfasından ilgili platformun binary'sini indirin ve `/usr/local/bin` (macOS/Linux) veya `Program Files` (Windows) klasörüne kopyalayın.

### 📦 Releases

En son sürümleri ve tüm platformlar için binary'leri [GitHub Releases](https://github.com/ozanbozkurtt/perf-analyzer/releases) sayfasından indirebilirsiniz.

| Platform | Binary |
|----------|--------|
| macOS (Apple Silicon M1/M2/M3) | `benchmark-macos-arm64` |
| macOS (Intel) | `benchmark-macos-x86_64` |
| Linux (ARM64) | `benchmark-linux-arm64` |
| Linux (x86_64) | `benchmark-linux-x86_64` |
| Windows (x86_64) | `benchmark-windows-x86_64.exe` |

### 🔨 Kaynak Koddan Derleme

Go 1.21 veya daha yeni bir sürümü varsa kaynaktan derleyebilirsiniz.

#### Ön Koşullar

- Go 1.21 veya daha yeni bir sürüm
- Git
- Linux, macOS veya Windows işletim sistemi

#### Adımlar

1. Depoyu klonlayın:
```bash
git clone https://github.com/ozanbozkurtt/perf-analyzer.git
cd perf-analyzer
```

2. Bağımlılıkları indirin:
```bash
go mod download
```

3. Aracı derleyin:
```bash
go build -o bin/benchmark cmd/benchmark/main.go
```

4. Multi-platform derlemesi (Linux, macOS, Windows için):
```bash
./build.sh v1.0.0
```

## Kullanım

### Temel Kullanım

#### Kurulmuş Binary'yle
```bash
benchmark
```

#### Kaynak Koddan
Derlediyseniz:
```bash
./bin/benchmark
```

Geliştirme sırasında doğrudan:
```bash
go run cmd/benchmark/main.go
```

### Komut Satırı Seçenekleri

```bash
benchmark [options]

Options:
  -h, --help      Bu yardımı göster
  -v, --version   Sürüm bilgisini göster
  -j, --json      Çıktıyı JSON formatında sun
```

### Çıktı Örneği

```
🔥 Hardware Performance Benchmark Tool

═══════════════════════════════════════════════════════════════
                    SİSTEM PERFORMANS RAPORU
═══════════════════════════════════════════════════════════════

📊 İŞLEMCİ (CPU) PERFORMANSI
├─ Çekirdek Sayısı: 8
├─ Mantıksal CPU: 16
├─ Model: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
└─ Max Frekans: 4800 MHz

🎮 RAM (Bellek) PERFORMANSI
├─ Toplam Bellek: 16 GB
├─ Kullanılan: 8.2 GB
├─ Boş: 7.8 GB
└─ Kullanım Yüzdesi: 51.25%

💿 DİSK PERFORMANSI
├─ Toplam Alan: 512 GB
├─ Kullanılan: 256 GB
├─ Boş: 256 GB
└─ Kullanım Yüzdesi: 50%

🌐 AĞ PERFORMANSI
├─ Bağlantı Durumu: Aktif
├─ Bağlantı Hızı: 1000 Mbps
└─ IP Adresi: 192.168.1.100

═══════════════════════════════════════════════════════════════
                       KALİTE PUANLANDIRMASI
═══════════════════════════════════════════════════════════════

[CPU] ████████░░ 8.2/10  (Mükemmel)
[RAM] ██████░░░░ 6.5/10  (İyi)
[Disk] ███████░░░ 7.0/10 (İyi)

Genel Skor: 7.2/10 ✓

═══════════════════════════════════════════════════════════════
                           ÖNERİLER
═══════════════════════════════════════════════════════════════

✓ CPU performansı mükemmel durumda
⚠ RAM kullanımı %50 üstünde - gereksiz uygulamaları kapatın
✓ Disk alanı uygun
```

## Proje Yapısı

```
perf-analyzer/
├── cmd/
│   └── benchmark/
│       └── main.go              # Ana giriş noktası
├── internal/
│   ├── benchmark/
│   │   └── runner.go            # Benchmark çalıştırıcı
│   ├── types/
│   │   └── types.go             # Veri yapıları
│   ├── cpu/
│   │   └── cpu.go               # CPU analizi
│   ├── memory/
│   │   └── memory.go            # Bellek analizi
│   ├── disk/
│   │   └── disk.go              # Disk analizi
│   ├── network/
│   │   └── network.go           # Ağ analizi
│   └── system/
│       └── system.go            # Sistem bilgileri
├── go.mod                        # Go modülü tanımı
├── go.sum                        # Bağımlılık kontrol dosyası
└── README.md                     # Bu dosya
```

## Teknik Detaylar

### Kullanılan Kütüphaneler

- **gopsutil**: Sistem bilgileri ve performans metrikleri (CPU, RAM, Disk, Ağ)
- **color**: Renkli terminal çıktısı

### Mimarisi

Proje temiz bir Go mimarisi takip eder:

1. **cmd/**: Komut satırı uygulamasının giriş noktası
2. **internal/**: Dış paketler tarafından erişilemeyen iç paketler
3. **types/**: Paylaşılan veri yapıları
4. **cpu, memory, disk, network**: Alan spesifik analiz modülleri
5. **benchmark**: Tüm analizleri koordine eden ana runner

## Geliştirme

### Testleri Çalıştırma

```bash
go test ./...
```

Race condition deteksiyonu ile:

```bash
go test -race ./...
```

Kodu biçimlendirme:

```bash
go fmt ./...
goimports -w ./...
```

Statik analiz:

```bash
go vet ./...
staticcheck ./...
```

## Katkıda Bulunma

Katkılarınız hoşlanır! Aşağıdaki adımları izleyin:

1. **Fork** edin: Repository'nin bir kopyasını oluşturun
2. **Branch** oluşturun: 
   ```bash
   git checkout -b feature/yeni-ozellik
   ```
3. **Değişiklikleri commit edin**:
   ```bash
   git commit -m "feat: yeni özellik açıklaması"
   ```
4. **Push** edin:
   ```bash
   git push origin feature/yeni-ozellik
   ```
5. **Pull Request** açın

### Geliştirme Kuralları

- Go idiomlarına uyun
- Tüm yeni özellikler için test yazın (%80+ kapsama hedefi)
- Kodu `gofmt` ve `goimports` ile biçimlendirin
- Anlamlı commit mesajları yazın
- Dökümentasyonu güncellyin

## Lisans

Bu proje MIT Lisansı altında dağıtılmaktadır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

## Yazarlar

- **Ozan Bozkurt** - Oluşturucu & Geliştirici

## İletişim

Sorular veya öneriler için GitHub Issues'ı kullanın veya ozan.bozkurt@ode.al adresine e-posta gönderin.

## Sürümler

### Mevcut Sürümleri Görmek

```bash
# Cihazınızda kurulu sürümü kontrol edin
benchmark --version

# Tüm sürümleri GitHub'dan görüntüleyin
# https://github.com/ozanbozkurtt/perf-analyzer/releases
```

### Güncelleme

Yeni sürüme güncellemek için kurulum scriptini tekrar çalıştırın. En son sürümü otomatik olarak indirecektir:

```bash
# macOS/Linux
curl -sL https://github.com/ozanbozkurtt/perf-analyzer/releases/download/latest/install.sh | bash

# Windows (PowerShell Admin)
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/ozanbozkurtt/perf-analyzer/main/install.ps1'))
```

## CI/CD İntegrasyonu

### GitHub Actions

Depo bir GitHub Actions workflow'u ile donatılmıştır. Her yeni tag oluşturulduğunda otomatik olarak multi-platform binary'ler derlenip release yapılır.

Tag oluşturmak için:
```bash
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

### GoReleaser (Opsiyonel)

Yerel olarak GoReleaser ile release yapabilirsiniz:

```bash
# GoReleaser yükleyin
brew install goreleaser  # macOS
# veya
# https://goreleaser.com/install/

# Release yapın (draft mod)
goreleaser release --skip=publish --skip=validate

# Production release
goreleaser release
```

## Destek

Aracı faydalı bulursanız, lütfen:
- ⭐ Repository'e star verin
- 🐛 [Sorun bildirin](https://github.com/ozanbozkurtt/perf-analyzer/issues)
- 💡 [Önerilerde bulunun](https://github.com/ozanbozkurtt/perf-analyzer/discussions)
- 🤝 Katkıda bulunun

---

**Mutlu benchmarking! 🚀**

## Sık Sorulan Sorular (FAQ)

### S: Go yüklü değilse kullanabilir miyim?
**C:** Evet! GitHub Releases sayfasından derlenmiş binary'leri indirebilirsiniz. Go kurulması gerekmez.

### S: Hangi platformlar destekleniyor?
**C:** macOS (Intel ve Apple Silicon), Linux (x86_64 ve ARM64) ve Windows (x86_64) desteklenmektedir.

### S: Binary'nin boyutu ne kadar?
**C:** Derlenmiş binary'ler 3-5 MB arasındadır. CGO_ENABLED=0 ile derlendikleri için herhangi bir bağımlılığa ihtiyaçları yoktur.

### S: Nasıl kontrol edebilirim ki binary güvenli?
**C:** SHA256 checksum'ları GitHub Releases sayfasında mevcuttur. İndirdikten sonra doğrulayabilirsiniz:
```bash
sha256sum benchmark-* > checksums.txt
sha256sum -c checksums.txt
```

### S: Windows'ta kurulum hakkında bilgi?
**C:** PowerShell scriptini kullanın veya binary'yi manuel olarak Program Files'a kopyalayın. Administrator izni gerekebilir.
