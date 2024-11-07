// 상태 코드 및 응답 바디를 저장하는 미들웨어
// http.ResponseWriter 인터페이스는 읽기 관련 메서드를 가지고 있지 않음
// 따라서 그대로 사용하면 응답 바디나 상태 코드를 이용할 수 없다.
// -> 로그에 출력하려면 wrapper 구조체를 사용해야한다.
// 아래 코드를 사용하면 http.ResponseWriter 인터페이스에서 얻을 수 없는 상태 코드나 바디 내용을 로그에 저장할 수 있게 된다.
// 응답이나 요청 바디 등에 개인정보가 있을 때, 이런 것을 관리하는 측면도 잘 살펴서 저장해야한다.
// 응답 내용을 후킹하는 함수

package exercise

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// rwWrapper 구조체는 http.ResponseWriter를 감싸는 구조체로, 상태 코드와 응답 바디를 저장하기 위한 필드를 추가
type rwWrapper struct {
	rw     http.ResponseWriter // 실제 응답을 처리할 http.ResponseWriter
	mw     io.Writer           // 응답을 여러 곳에 동시에 작성하기 위한 io.Writer
	status int                 // 상태 코드를 저장
}

// NewRwWrapper 함수는 응답 바디를 로그에 기록할 수 있도록 io.MultiWriter를 통해
// http.ResponseWriter와 io.Writer를 동시에 기록하는 구조체를 반환
func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw: rw,
		mw: io.MultiWriter(rw, buf), // 응답을 실제 응답 스트림과 버퍼에 동시에 기록
	}
}

// Header 메서드는 http.ResponseWriter의 Header 메서드를 호출해 헤더를 반환
func (r *rwWrapper) Header() http.Header {
	return r.rw.Header()
}

// Write 메서드는 응답을 http.ResponseWriter와 버퍼에 동시에 기록하고,
// 처음 Write 메서드가 호출될 때 기본 상태 코드를 200(OK)으로 설정
func (r *rwWrapper) Write(i []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.mw.Write(i)
}

// WriteHeader 메서드는 상태 코드를 설정하고 이를 실제 응답에도 반영
func (r *rwWrapper) WriteHeader(statusCode int) {
	r.status = statusCode           // 상태 코드 저장
	r.rw.WriteHeader(statusCode)     // 실제 응답에 상태 코드 설정
}

// NewLogger 함수는 로깅 미들웨어를 반환하며, 이 미들웨어는 HTTP 요청을 처리하고
// 응답의 상태 코드와 바디를 로그에 기록
func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 버퍼를 생성하여 응답 바디를 저장할 준비
			buf := &bytes.Buffer{}
			// rwWrapper로 http.ResponseWriter 감싸기
			rww := NewRwWrapper(w, buf)
			// 다음 핸들러로 요청 전달
			next.ServeHTTP(rww, r)
			// 버퍼에 저장된 응답 바디를 로그에 출력
			l.Printf("Response Body: %s", buf.String())
			// 응답 상태 코드 출력
			l.Printf("Status Code: %d", rww.status)
		})
	}
}
