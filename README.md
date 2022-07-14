# Pemantau Resi

> Catatan: **kode ini hanya sebagai pembelajaran semata.**

## Apa ini?

Pemantau resi sederhana, setiap kali ada perubahan dari riwayat resi, maka bot akan mengirimkannya via chat ke telegram kamu.

Script ini ditulis dengan cepat, sehingga kode kurang tertata rapi. Jika kamu ingin membantu merapikan atau membetulkan bug, silakan membuat pull request, saya akan sangat senang.

## Bagaimana cara kerjanya?

1. Script akan membaca list resi dari [resi.json](/resi.json).

2. Script akan mencsrape web [resi.id](https://resi.id) dengan menggunakan link resi yang ada.

3. Untuk pengulangan pertama, script akan menyimpan riwayat resi yang sudah discrape, dalam bentuk json, ke dalam folder [hisotry](/history/), dan mengirimkan update ke telegram menggunakan bot([bot.go](/bot.go)).

4. Pada pengulangan selanjutnya, hasil scrape yang terbaru akan **dibandingkan** dengan data scrape yang sebelumnya sudah tersimpan.

   - Jike kedua data **sama**, berarti **tidak ada perubahan**, script akan berhenti.
   - Jika kedua data **berbeda**, berarti telah terdapat **update data.** Maka bot akan mengirimkan riwayat resi terbaru ke telegram. Kemudian mensave riwayat resi dengan versi terbaru ke github.

5. Logic tersebut akan dijalankan setiap 15 menit sekali menggunakan Github Action ([.github/workflows/update.yml](/.github/workflows/update.yml)).

## Bagaimana cara menggunakannya?

> Baca perlahan, sabar.

1. Fork repository ini, klik tombol [Fork](https://github.com/zakiego/pantau-resi/fork) di sebelah kanan atas.

2. Edit file [config.json](/config.json)

   - Pada bagian `chat_id`, masukkan ID telegram-mu. Untuk mengecek ID telegram, bisa menggunakan bot [RawDataBot](https://t.me/raw_data_bot). Chat ke bot tersebut, lalu kemudian bot-nya akan menampilkan ID telegram-mu.

   - Pada bagian `bot_token`, masukkan token dari botnya. Untuk bagian ini, kamu harus membuat bot telegram terlebih dahulu. Baca artikel singkat ini untuk membuat bot telegram: [Cara Membuat BOT Telegram, Tak Sampai 5 Menit Jadi! Mudah dan Simpel!](https://kumparan.com/berita-terkini/cara-membuat-bot-telegram-tak-sampai-5-menit-jadi-mudah-dan-simpel-1v3iKFA8Jkt/1) - Kumparan

3. Edit file [resi.json](/resi.json)

   Saya menggunakan web [resi.id](https://resi.id/) sebagai sumber data.

   Masukkan resi yang kamu miliki di web tersebut, klik bagikan, lalu klik icon link (warna oranye, ujung kanan), salin linknya (misal: https://resi.id/s/123456788).

   Kemudian edit file [resi.json](/resi.json), masukkan link tadi. Sehingga berbentuk seperti ini:

   ```json
   {
     "resi": [
       {
         "link": "https://resi.id/s/123456788"
       },
       {
         "link": "https://resi.id/s/8738474"
       },
       {
         "link": "https://resi.id/s/36474837"
       }
     ]
   }
   ```

4. Selesai. Github action akan menjalankan script [.github/workflows/update.yml](/.github/workflows/update.yml) setiap 15 menit.

   Untuk mengubah rentang waktu pengulangan, edit pada baris:

   ```yml
   on:
     schedule:
       - cron: "*/15 * * * *" # every 15th minute
   ```

   Pada bagain `cron` ganti dengan durasi yang diinginkan.
   Klik https://crontab.guru/examples.html untuk melihat referensi waktu yang bisa digunakan.

## Dibuat Dengan

- [Golang](https://go.dev/) - Bahasa pemrograman yang digunakan (kamu bisa menggunakan bahasa apa pun, yang penting logicnya dapet)
- [goquery](https://github.com/PuerkitoBio/goquery) - Library untuk parse html di Golang
- [Github Action](https://github.com/features/actions) - Untuk menjalankan script
