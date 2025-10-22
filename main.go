package main

import (
	"fmt"
	"net/url"

	_ "github.com/udistrital/agora_api_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerrorv2"
	"github.com/udistrital/utils_oas/ssm"
	"github.com/udistrital/utils_oas/xray"
)

func init() {
	if beego.AppConfig.String("parameterStore") != "" {
		parameterStore := "/" + beego.AppConfig.String("parameterStore") +
			"/" + beego.AppConfig.String("appname") + "/db/"

		username, err := ssm.GetParameterFromParameterStore(parameterStore + "username")
		if err != nil {
			logs.Critical("Error retrieving username: %v", err)
		}

		err = beego.AppConfig.Set("PGuser", username)
		if err != nil {
			logs.Critical("Failed to set PGuser env var: %v", err)
		}

		password, err := ssm.GetParameterFromParameterStore(parameterStore + "password")
		if err != nil {
			logs.Critical("Error retrieving password: %v", err)
		}

		err = beego.AppConfig.Set("PGpass", password)
		if err != nil {
			logs.Critical("Failed to set PGpass: %v", err)
		}
	}

	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+url.QueryEscape(beego.AppConfig.String("PGpass"))+"@"+beego.AppConfig.String("PGhost")+":"+beego.AppConfig.String("PGport")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschema"))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	xray.InitXRay()
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	beego.InsertFilter("*", beego.BeforeExec, SecurityHeaders)
	beego.Run()
}

func SecurityHeaders(ctx *context.Context) {
	fmt.Println("headers set")
	ctx.Output.Header("Clear-Site-Data", "'cache', 'cookies', 'storage', 'executionContexts'")
	ctx.Output.Header("Cross-Origin-Embedder-Policy", "require-corp")
	ctx.Output.Header("Cross-Origin-Opener-Policy", "same-origin")
	ctx.Output.Header("Cross-Origin-Resource-Policy", "same-origin")
	ctx.Output.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
	ctx.Output.Header("Referrer-Policy", "no-referrer")
	ctx.Output.Header("Server", "")
	ctx.Output.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	ctx.Output.Header("X-Content-Type-Options", "nosniff")
	ctx.Output.Header("X-Frame-Options", "DENY")
	ctx.Output.Header("X-Permitted-Cross-Domain-Policies", "none")
}
