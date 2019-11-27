package util

import "testing"

func TestFignature(t *testing.T) {
	sign := Signature("1574867028", "1761970467", "aes", "44500f2b4b10dd95588559dae93be5f259f4f68c", "xiaoenai")
	t.Logf("Sign-> %s", sign)
}
