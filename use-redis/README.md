## 使用redis，存取键值对
需要在本机上安装好docker
`docker run --rm -itd -p 6379:6379 --name redis redis:latest`

`go get github.com/garyburd/redigo/redis`

`cd ${GOPATH}/src`

`git clone https://github.com/fwhezfwhez/go-learn.git`

`cd ${GOPATH}/src/go-learn`

`cd use-redis`

`go run main.go`
