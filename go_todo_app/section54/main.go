// 요청을 받아서 응답 메세시 생성 서버
// 포트 번호는 18080을 고정해서 서버 실행

// http.ListenAndServe:
// ListenAndServe는 지정된 포트(:18080)에서 HTTP 서버를 시작하고, 인수로 전달된 핸들러(http.HandlerFunc)를 설정합니다.
// 여기서는 익명 함수를 http.HandlerFunc로 전달하여, 요청을 받을 때마다 이 함수가 실행되도록 합니다.

// 핸들러 함수:
// 익명 함수 func(w http.ResponseWriter, r *http.Request)는 모든 요청을 처리합니다.
// r.URL.Path[1:]을 통해 URL 경로를 가져오고, 이를 응답 메시지에 포함시켜 클라이언트에 반환합니다. 예를 들어, http://localhost:18080/world 요청이 오면 Hello, world!가 응답됩니다.

// 에러 처리:
// ListenAndServe 함수는 에러를 반환할 수 있으며, 에러가 발생하면 이를 로그에 출력하고 프로그램을 종료합니다.

package section54

import(
	"fmt"
	"net/http"
	"os"
)

func Main() {
	err := http.ListenAndServe(
		":18080", // 포트 18080에 서버를 바인딩
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// URL 경로를 이용해 응답 메시지 생성
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]) 
		}),
	)
	if err != nil {
		// 서버 오류 메시지 출력
		fmt.Printf("failed to terminate server : %v", err) 
		os.Exit(1) // 비정상 종료
	}
}
