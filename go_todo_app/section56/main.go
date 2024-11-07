package section56

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"golang.org/x/sync/errgroup"
)
func main(){
	if len(os.Args) != 2{
		log.Printf("need port number\n")
		os.Exit(1)
	}
	p := os.Args[1]
	l,err := net.Listen("tcp", ":"+p)
	if err != nil{
		log.Fatalf("failed to listen port %s: %v", p, err)
	}
	if err := run(context.Background(), l); err != nil{
		log.Printf("failed to terminate server : %v", err)
		os.Exit(1)
	}
}


// main 함수 정의는 생략
func run(ctx context.Context, l net.Listener) error{
	s := &http.Server{
		// Addr: ":18080", -- 빠진 코드(포토를 자동으로 할당받기 위해)
		// 인수로 받은 net.Listener 를 이용하므로 Addr 필드는 지정하지 않는다.
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 HTTP 서버를 실행한다.
	eg.Go(func() error{
		// ListenAndServe 메서드가 아닌 Serve 메서드로 변경한다.
		// if err := s.ListenAndServe(); err != nil &&
		// 빠진 줄 -- 무슨 포트를 할당받는지 모르기 때문에 인자로 받아옴
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed는
			// http http.Server.Shutdown()가 정상 종료됐다고 표시하므로 문제없다.
		err != http.ErrServerClosed{
		log.Printf("failed to close : %+v", err)
		return err
		}
		return nil
	})
	// 채널로부터의 알림(종료 알림)을 기다린다.
	<- ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil{
		log.Printf("failed to shutdown : %+v", err)
	}
	// Go 메서드로 실행한 다른 고루틴 종료를 기다린다.
	return eg.Wait()
}