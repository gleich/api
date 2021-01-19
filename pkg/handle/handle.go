package handle

import (
	"context"
	"net/http"
	"time"

	"github.com/Matt-Gleich/lumber"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
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
					lumber.Info("Responded to request that had errors")
				}
				lumber.Success("Responded to request")
			}
		},
	})

	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour}), h.ServeHTTP))
	err := http.ListenAndServe(":8080", nil)
	lumber.Fatal(err, "Failed to listen and serve for requests")
}