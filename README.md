
# Simple API 

simple api ada aplikasi rest api sederhana dengan menggunakan bahasa golang. 

## Fitur local

- Register
- Login User
- Edit User
- Mencari user menggunakan nama 
## Fitur Tambahan
- Consume public API dari 
```
https://ddragon.leagueoflegends.com/cdn/6.24.1/data/en_US/champion.json
```
dengan enpoint
```
/search/:name
```
## Open API

Untuk Open API bisa lihat selengkapnya [disini](https://app.swaggerhub.com/apis-docs/satriacening/simpeApi/1.0.0)


## Menjalankan Lokal

Kloning project

```bash
  $ git clone git@github.com:ALTA-BE7-SATRIA-PUTRA/group_project3.git
```

Masuk ke direktori project

```bash
  $ cd ~/project
```
Buat `database` baru

Buat sebuah file dengan nama di dalam folder root project `.env` dengan format dibawah ini. Sesuaikan configurasi di komputer lokal

```bash
  export APP_PORT=""
  export JWT_SECRET=""
  export DB_PORT=""
  export DB_DRIVER=""
  export DB_NAME=""
  export DB_ADDRESS=""
  export DB_USERNAME=""
  export DB_PASSWORD=""
```

Jalankan aplikasi 

```bash
  $ go run main.go
```


## Authors

- [@satriacening](https://github.com/satriacening)

 
