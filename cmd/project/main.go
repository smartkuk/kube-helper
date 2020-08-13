// go.mod 프로젝트를 $GOPATH 기준의 디렉토리에 옮겨주는 모듈
// go build -o bin/project/kube-project cmd/project/main.go
// go run ./cmd/project/main.go

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/smartkuk/kube-helper/pkg/file"
)

// EnvGopath GO 패스 환경변수 이름
const EnvGopath = "GOPATH"

// EnvModule GO 패스 하위 모듈 환경변수 이름
const EnvModule = "MODULE"

func main() {
	p, ok := os.LookupEnv(EnvGopath)
	if !ok {
		panic("Fail to look up GOPATH env variable.")
	}

	log.Printf("GOPATH: %s, exist result: %t", p, ok)
	m, ok := os.LookupEnv(EnvModule)
	if !ok {
		panic("Fail to look up MODULE env variable.")
	}
	log.Printf("MODULE: %s, exist result: %t", m, ok)

	// 현재 작업디렉토리 획득
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	log.Println(dir)
	log.Println(file.GetBaseName(dir))
	baseName := file.GetBaseName(dir)
	dest := filepath.Join(p, m, baseName)
	log.Printf("destination: %s", dest)
	err = file.CreateIfNotExists(dest, 0755)
	if err != nil {
		panic(err)
	}
	err = file.CopyDirectory(dir, dest)
	if err != nil {
		panic(err)
	}
}
