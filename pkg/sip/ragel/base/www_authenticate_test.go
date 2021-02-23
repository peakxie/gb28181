package base_test

import (
	"testing"

	"github.com/peakxie/gb28181/pkg/sip/ragel/base"
)

const (
	www_auth = " Digest username=\"34020000001320300001\", realm=\"172.21.1.9\", nonce=\"ae491f82da1cd66dad1b1bed3c386cd25a548ce34e038aeaeb4023ed0dcd67d5\", uri=\"sip:34020000002000000002@172.21.1.9:15060\", response=\"adb4e76d5057d162e5902bad47760323\", algorithm=MD5, opaque=\"994dd075d0572c63c3d662e23d7ae31dea4fa8dda096d2c1174a6a7308fba476\""
)

//ragel -Z -G2 -o www_authenticate.go www_authenticate.rl
func TestParseWwwAuth(t *testing.T) {

	params, err := base.ParseWwwAuth([]byte(www_auth))
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	params.M.Range(func(key interface{}, value interface{}) bool {
		t.Logf("%v:%v", key, value)
		return true
	})
}

func BenchmarkParseWwwAuth(b *testing.B) { // 26653 ns/op
	for i := 0; i < b.N; i++ {
		base.ParseWwwAuth([]byte(www_auth))
	}
}
