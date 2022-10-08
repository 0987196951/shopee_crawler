module shopee.rd/cmd

go 1.19

replace shopee.rd/utils => ../utils

replace shopee.rd/database => ../database

require shopee.rd/utils v0.0.0-00010101000000-000000000000

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	shopee.rd/config v0.0.0-00010101000000-000000000000 // indirect
)

replace shopee.rd/crawler => ../crawler

replace shopee.rd/config => ../config
