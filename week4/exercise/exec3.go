// 복원 미들웨어
package exercise

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// `defer`를 사용하여 이 함수가 완료된 후 특정 동작을 수행
		defer func() {
			// `recover()`는 `panic`이 발생했을 때 이를 복구하고 `panic`의 값을 반환함
			if err := recover(); err != nil {
				// 오류 메시지를 JSON 형식으로 변환
				jsonBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprintf("%v", err),
				})

				// 헤더에 JSON 응답 형식 설정
				w.Header().Set("Content-Type", "application/json")
				// HTTP 상태 코드를 500(서버 오류)로 설정
				w.WriteHeader(http.StatusInternalServerError)
				// 변환된 JSON 오류 메시지를 응답으로 보냄
				w.Write(jsonBody)
			}
		}()
		
		// 패닉이 발생하지 않을 경우 `next` 핸들러 호출
		next.ServeHTTP(w, r)
	})
}
