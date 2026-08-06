package agent

import (
	"os"
	"testing"
)

func TestApplyConfiguredHTTPProxy(t *testing.T) {
	t.Setenv("ZOT_HOME", t.TempDir())
	for _, key := range []string{"HTTP_PROXY", "HTTPS_PROXY", "http_proxy", "https_proxy"} {
		t.Setenv(key, "")
	}
	if err := SaveConfig(Config{HTTPProxy: "  http://proxy.example:8080  "}); err != nil {
		t.Fatal(err)
	}

	applyConfiguredHTTPProxy()

	if got := getenvForTest("HTTP_PROXY"); got != "http://proxy.example:8080" {
		t.Fatalf("HTTP_PROXY = %q", got)
	}
	if got := getenvForTest("HTTPS_PROXY"); got != "http://proxy.example:8080" {
		t.Fatalf("HTTPS_PROXY = %q", got)
	}
}

func TestApplyConfiguredHTTPProxyPreservesEnvironment(t *testing.T) {
	t.Setenv("ZOT_HOME", t.TempDir())
	t.Setenv("HTTP_PROXY", "http://http-env.example:8080")
	t.Setenv("http_proxy", "")
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("https_proxy", "http://https-env.example:8080")
	if err := SaveConfig(Config{HTTPProxy: "http://config.example:8080"}); err != nil {
		t.Fatal(err)
	}

	applyConfiguredHTTPProxy()

	if got := getenvForTest("HTTP_PROXY"); got != "http://http-env.example:8080" {
		t.Fatalf("HTTP_PROXY = %q", got)
	}
	if got := getenvForTest("HTTPS_PROXY"); got != "" {
		t.Fatalf("HTTPS_PROXY = %q, want empty because lowercase environment setting takes precedence", got)
	}
	if got := getenvForTest("https_proxy"); got != "http://https-env.example:8080" {
		t.Fatalf("https_proxy = %q", got)
	}
}

func TestApplyConfiguredHTTPProxyIgnoresEmptySetting(t *testing.T) {
	t.Setenv("ZOT_HOME", t.TempDir())
	for _, key := range []string{"HTTP_PROXY", "HTTPS_PROXY", "http_proxy", "https_proxy"} {
		t.Setenv(key, "")
	}
	if err := SaveConfig(Config{HTTPProxy: "  "}); err != nil {
		t.Fatal(err)
	}

	applyConfiguredHTTPProxy()

	if got := getenvForTest("HTTP_PROXY"); got != "" {
		t.Fatalf("HTTP_PROXY = %q", got)
	}
	if got := getenvForTest("HTTPS_PROXY"); got != "" {
		t.Fatalf("HTTPS_PROXY = %q", got)
	}
}

func getenvForTest(key string) string {
	return os.Getenv(key)
}
