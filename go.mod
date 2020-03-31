module github.com/weekndCN/shorturl

go 1.14


require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.55.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
)


replace (
	github.com/weekndCN/shorturl/conf => /rwplus/rwplus-shorturl/conf
	github.com/weekndCN/shorturl/models => /rwplus/rwplus-shorturl/models
	github.com/weekndCN/shorturl/pkg => /rwplus/rwplus-shorturl/pkg
	github.com/weekndCN/shorturl/pkg/e => /rwplus/rwplus-shorturl/pkg/e
	github.com/weekndCN/shorturl/routers => /rwplus/rwplus-shorturl/routers
	github.com/weekndCN/shorturl/routers/api => /rwplus/rwplus-shorturl/routers/api
)
