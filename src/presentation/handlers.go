package presentation

import (
	"cuturl/src/application"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func MakeUrlHandler(ioc InteractorFactory) *router.Router {
	httpRouter := router.New()
	httpRouter.GET("/", indexHandler)
	httpRouter.POST("/shorten", func(ctx *fasthttp.RequestCtx) {
		makeUrlHandler(ctx, ioc)
	})
	httpRouter.GET("/{code}", func(ctx *fasthttp.RequestCtx) {
		urlHandler(ctx, ioc)
	})
	return httpRouter
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetBodyString("Hello, indexHandler!\n")
}

func makeUrlHandler(ctx *fasthttp.RequestCtx, ioc InteractorFactory) {
	url := ctx.FormValue("url")

	makeUrl := ioc.MakeUrl()
	response, err := makeUrl.Execute(&application.MakeUrlRequest{Url: string(url)})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.SetBodyString("{\"error\": \"invalid url\"}")
		return
	}

	ctx.SetContentType("application/json")
	responseJson, err := json.Marshal(response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("{\"error\": \"internal error\"}")
		return
	}

	ctx.SetBodyString(string(responseJson))
}

func urlHandler(ctx *fasthttp.RequestCtx, ioc InteractorFactory) {
	code := ctx.UserValue("code").(string)

	getUrl := ioc.GetUrl()
	response, err := getUrl.Execute(&application.GetUrlRequest{Code: code})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.SetBodyString("{\"error\": \"not found\"}")
		return
	}

	ctx.Redirect(response.Url, fasthttp.StatusMovedPermanently)
}
