package handle

import (
	"context"
	"net/http"

	"github.com/gleich/lumber"
	"github.com/didip/tollbooth"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Run the handler
func Run(s graphql.Schema) {
	h := handler.New(&handler.Config{
		Schema:     &s,
		Playground: true,
		ResultCallbackFn: func(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {
			if ctx.Err() != nil {
				lumber.Error(ctx.Err(), "Failed to respond to request")
			} else {
				if result.HasErrors() {
					lumber.Warning("Request had errors", result.Errors)
				}
				lumber.Success("Responded to request")
			}
		},
	})

	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(10, nil), h.ServeHTTP))
	err := http.ListenAndServe(":80", nil)
	lumber.Fatal(err, "Failed to listen and serve for requests")
}
