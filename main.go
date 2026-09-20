// storefront — the smallest service that can still be broken on purpose.
//
// It serves one line of text, taken from the BANNER environment variable, on
// :5678. The port matches hashicorp/http-echo deliberately: the manifests in
// argocd-class-resources already target 5678, so swapping this image in for the
// prebuilt one changes nothing but the image reference.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const addr = ":5678"

// Banner is what the service answers with. It is read once, at process start,
// which is the behaviour S01 L04 depends on: change the ConfigMap and nothing
// moves until the pod restarts.
func Banner() string {
	if b := os.Getenv("BANNER"); b != "" {
		return b
	}
	return "storefront — no banner set"
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, Banner())
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("storefront listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
