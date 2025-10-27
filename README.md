# exact-time
This go module allows you to get the exact current time using the NTP protocol.

## Использование

Для того чтобы получить актуальное время, достаточно запустить программу через:

```bash
go run main.go
```

Также вы можете поменять адрес NTP сервера. Это возможно сделать через `.env` файл:

```bash
NTPSERVER="<ваш_адрес_ntp_сервера>"
```

Если не указать адрес NTP сервера, он будет выбран по умолчанию: `0.beevik-ntp.pool.ntp.org`


