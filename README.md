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

### Ön Koşullar

- Go 1.21 veya daha yeni bir sürüm
- Linux, macOS veya Windows işletim sistemi

### Adımlar

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

## Kullanım

### Temel Kullanım

Aracı çalıştırmak çok basit:

```bash
./bin/benchmark
```

veya geliştirme sırasında doğrudan:

```bash
go run cmd/benchmark/main.go
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

## Destek

Aracı faydalı bulursanız, lütfen:
- ⭐ Repository'e star verin
- 🐛 Sorun bildirin
- 💡 Önerilerde bulunun
- 🤝 Katkıda bulunun

---

**Mutlu benchmarking! 🚀**
