package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kubermatic/api/handler"
	"github.com/kubermatic/api/provider/cloud"
	"github.com/kubermatic/api/provider/kubernetes"
	"golang.org/x/net/context"

	ghandlers "github.com/gorilla/handlers"
)

func main() {
	// parse flags
	kubeconfig := flag.String("kubeconfig", ".kubeconfig", "The kubeconfig file path with one context per Kubernetes provider")
	auth := flag.Bool("auth", true, "Activate authentication with JSON Web Tokens")
	jwtKey := flag.String("jwt-key", "", "The JSON Web Token validation key, encoded in base64")
	flag.Parse()

	// create CloudProviders
	cps := cloud.Providers()
	// create KubernetesProvider for each context in the kubeconfig
	kps, err := kubernetes.Providers(*kubeconfig, cps)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	ctx := context.Background()
	b := handler.NewBinding(ctx, kps, cps, *auth, *jwtKey)
	mux := mux.NewRouter()
	b.Register(mux)

	log.Fatal(http.ListenAndServe(":8080", ghandlers.CombinedLoggingHandler(os.Stdout, mux)))
}
