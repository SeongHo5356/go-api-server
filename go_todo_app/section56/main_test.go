// 서버를 실행한 상태에서 포트가 겹치므로 테스트 불가
// 현재 코드의 문제점
// 1. 테스트 완료 후에 종료할 방법이 없다
// 2. 출력을 검증하기 어렵다
// 3. 이상 처리 시에 os.Exit 함수가 호출된다.
// 4. 포트 번호가 고정돼 있어서 테스트에서 서버가 실행되지 않을 수 있다.
// -> main 함수에서 처리를 분리해서 run 함수로 옮기도록 하여 해결
// -> 웹 서버 구현 뿐 아닌, 명령줄 구현에도 사용할 수 있다.

package section56

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T){

	// -- 바뀐 코드 --
	l, err := net.Listen("tcp", "localhost:0" )
	if err != nil{
		t.Fatalf("failed to listen port %v", err)
	}
	// -- 바뀐 코드 --

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, l)
	})

	in := "message"
	// rsp, err := http.Get("http://localhost:18080/" + in)
	// -- 바뀐 코드 -- 
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// 어떤 포트 번호를 리슨하고 있는지 확인
	t.Logf("try request to %q", url)
	rsp, err := http.Get(url)
	// -- 바뀐 코드 --

	if err != nil{
		// %+v는 오류의 전체 스택 트레이스와 상세 정보를 출력
		// 특히 errors 패키지를 활용해 감싸진 오류(fmt.Errorf나 errors.Wrap 등)를 출력할 때 유용
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil{
		// %v는 기본적인 문자열 정보만 출력 
		// 여기선 오류의 간단한 설명만 필요하다고 생각할 때 %v를 주로 사용
		t.Fatalf("failed to read body: %v", err)
	}

	// Http 서버의 반환값을 검증
	// Sprintf() 용도: 주로 형식문자를 포함한 전체 문자열을 변수에 저장하고 출력할 때 사용`	`
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run 함수에 종료 알림을 전송한다.
	cancel()

	// run 함수의 반환값을 검증한다.
	// eg.Wait()의 결과를 err 변수에 할당
	// err가 nil이 아닌지(즉, 오류가 발생했는지)를 확인
	// err가 nil이 아니라면, t.Fatal(err)로 테스트를 실패하게 하고, 에러 내용을 출력
	if err := eg.Wait(); err != nil{
		t. Fatal(err)
	}

	// go test -v ./...를 통해서 실행
	// -v
	// -v 는 verbose 모드를 의미
	// 테스트를 실행할 때 더 많은 정보를 출력하도록 한다.
	// 테스트 함수의 이름과 함께 각 테스트의 성공 또는 실패 결과를 한눈에 보여줌
	// ./...
	// 현재 디렉토리와 모든 하위 디렉토리에 있는 모든 패키지를 대상으로 테스트를 실행
	// ./ 현재 디렉토리를 가리키고, ... 는 그 하위의 모든 패키지를 포함한다.
	// 현재 디렉토리에 여러 개의 서브 패키지가 있는 경우, 모든 서브 패키지 재귀적으로 검색
}