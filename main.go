package main

import (
	"members/db"
	"members/server"
)

// @title members-api
// @version バージョン(1.0)
// @description 仕様書に関する内容説明
// @termsOfService 仕様書使用する際の注意事項

// @contact.name APIサポーター
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	db.Init()
	server.Init()
}
