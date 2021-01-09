package main

import (
	"members/db"
	"members/server"
)

// @title members-api
// @version バージョン(1.0)
// @description メンバーズ用API
// @termsOfService 仕様書使用する際の注意事項

// @contact.name tochika.biz@gmial.com
// @contact.url https://github.com/hiwatee
// @contact.email tochika.biz@gmial.com

// @license.name apache2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
func main() {
	db.Init()
	server.Init()
}
