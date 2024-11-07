// 요청 바디를 로그에 남기는 미들웨어
// Go의 HTTP 요청(*http.Request)는 스트림 구조로 한번 밖에 읽을 수 없다.
// 따라서 미들웨어 구현 내에서 요청 바디를 읽으면 후속 미들웨어나 HTTP 핸들러 처리 내에서 요청을 읽을 수 없다.
// 요청 body를 사용하려면 추가적으로 처리가 필요함

package exercise

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"go.uber.org/zap"
)

// 요청 바디를 로그에 남기고 다시 읽을 수 있도록 복원하는 미들웨어
func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청 바디 읽기
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to log request body: %v", zap.Error(err))
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 로그에 요청 바디 출력
		log.Printf("Request Body: %s", string(body))

		// 요청 바디를 다른 핸들러에서도 읽을 수 있도록 복원
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// 다음 핸들러 호출
		next.ServeHTTP(w, r)
	})
}
