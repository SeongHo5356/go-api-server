// 추가 정보를 사용한 미들웨어 패턴 구현 
// XXX 타입값을 인수로 전달할 수 없는 경우에 미들웨어 패턴을 반환하는 함수를 구현
package exercise

import "net/http"
// VersionAdder 함수는 다음과 같이 사용할 수 있는 미들웨어 함수를 반환한다.
// vmw := VersionAdder("1.0.1")
// http.Handle("/users", vmw(userHandler))

// AppVersion 을 string 타입으로 지정
type AppVersion string

func VersionAdder(v AppVersion) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			r.Header.Add("App-Version", string(v))
			next.ServeHTTP(w, r)
		})	
	}
}