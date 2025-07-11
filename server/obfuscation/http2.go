package obfuscation

import (
	"fmt"
	"io"
	"net"
)

func FakeHTTP2Handshake(conn net.Conn) error {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return err
  }
  data := buf[:n]
	// بررسی ساده درخواست HTTP (مثلاً POST /connect)
	if !isHTTPRequest(data) {
		return fmt.Errorf("not valid HTTP2-like handshake")
	}
	// می‌تونیم لاگ بزنیم، session ثبت کنیم و ... 

	// پاسخ ساده
	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"))
	return nil
}

func isHTTPRequest(data []byte) bool {
	return len(data) > 4 && (string(data[:4]) == "POST" || string(data[:3]) == "GET")
}
