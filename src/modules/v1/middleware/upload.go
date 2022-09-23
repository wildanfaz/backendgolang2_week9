package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if isErr := r.ParseMultipartForm(20); isErr != nil {
		libs.Response(nil, 400, "failed parse form", isErr).Send(w)
		return
	}

	file, handler, err := r.FormFile("image")

	defer file.Close()

	if err != nil {
		libs.Response(nil, 400, "failed to upload file", err).Send(w)
		return
	}

	//**check
	fmt.Println(handler.Filename, handler.Header, handler.Size)

	name := strings.ReplaceAll(time.Now().Format(time.UnixDate), ":", "-") + handler.Filename
	result, errs := os.Create("images/" + name)

	if errs != nil {
		libs.Response(nil, 400, "failed create file", errs).Send(w)
		return
	}

	if _, error := io.Copy(result, file); error != nil {
		libs.Response(nil, 400, "failed copy file", errs).Send(w)
		return
	}

	libs.Response(nil, 200, "success upload file", nil).Send(w)
}
