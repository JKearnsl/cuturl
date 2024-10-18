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
	ctx.SetContentType("text/html")
	ctx.SetBodyString(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>CutURL</title>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<link rel="stylesheet" href="https://unpkg.com/sakura.css/css/sakura.css" type="text/css">
		</head>
		<body>
		
		<h1>Cut your URL</h1>
		<p>This program allows you to shorten links</p>
		
		<label>
			<input alt="Input url" placeholder="https://google.com">
			<button>Shorten</button>
		</label>
		
		
		<script>
	
			const input = document.querySelector('input');
			const button = document.querySelector('button');
	
			button.addEventListener('click', () => {
				const url = input.value;
				formdata = new FormData();
				formdata.append("url", url);
				fetch('/shorten', {
					method: 'POST',
					body: formdata
				})
				.then(response => response.json())
				.then(data => {
					if (data.error) {
						input.value = data.error
					} else {
						input.value = input.value = window.location.origin + "/" + data.Code;
					}
				});
			});
		</script>
		
		</body>
		</html>
	`)
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
