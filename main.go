package main
import "net/http"
import "io/ioutil"

type NotFoundRedirectRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func (w *NotFoundRedirectRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	if status != http.StatusNotFound {
			w.ResponseWriter.WriteHeader(status)
	}
}

func (w *NotFoundRedirectRespWr) Write(p []byte) (int, error) {
	if w.status != http.StatusNotFound {
			return w.ResponseWriter.Write(p)
	}
	return len(p), nil
}

func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			nfrw := &NotFoundRedirectRespWr{ResponseWriter: w}
			h.ServeHTTP(nfrw, r)
			if nfrw.status >= 400 {
					f, err := ioutil.ReadFile("./public/404.html")
					if err != nil {
							http.Redirect(w, r, "/index.html", http.StatusFound);
							return;
					}
					w.Header().Set("Content-type", "text/html");
					w.Write(f);
			}
	}
}

func main() {
	http.HandleFunc("/", wrapHandler(http.FileServer(http.Dir("public"))));
	http.ListenAndServe(":8001", nil);
}