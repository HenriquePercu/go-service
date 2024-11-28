package web

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

// Handler HandlerFunc represents a function that handles a http request within our own
// little mini framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidHandler
}

func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}
func (a *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}

		ctx := setValues(r.Context(), v)

		if err := handler(ctx, w, r); err != nil {
			fmt.Println(err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
