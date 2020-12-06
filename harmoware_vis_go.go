package harmwoware_vis_go

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	gosocketio "github.com/mtfelian/golang-socketio"
)

type HarmowareVisGo struct{

}

func NewHarmowareVisGo() *HarmowareVisGo {
    return &HarmowareVisGo{

	}
}

func (hv *HarmowareVisGo) RunServer(address string) {
	ioserv := hv.runIOServer()
	log.Printf("Running Sio Server..\n")
	if ioserv == nil {
		os.Exit(1)
	}
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", ioserv)
	serveMux.HandleFunc("/", hv.assetsFileHandler)
	log.Printf("Starting Harmoware VIS  Provider on %s", address)
	err := http.ListenAndServe(address, serveMux)
	if err != nil {
		log.Fatal(err)
	}
}

func (hv *HarmowareVisGo) runIOServer()*gosocketio.Server {
	

	assetsDir := hv.getAssetsDir()
	log.Println("AssetDir:", assetsDir)

	server := gosocketio.NewServer()

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("Connected from %s as %s", c.IP(), c.Id())
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("Disconnected from %s as %s", c.IP(), c.Id())
	})

	return server
}


// assetsFileHandler for static Data
func (hv *HarmowareVisGo) assetsFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return
	}

	file := r.URL.Path
	//	log.Printf("Open File '%s'",file)
	if file == "/" {
		file = "/index.html"
	}
	assetsDir := hv.getAssetsDir()
	f, err := assetsDir.Open(file)
	if err != nil {
		log.Printf("can't open file %s: %v\n", file, err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Printf("can't open file %s: %v\n", file, err)
		return
	}
	http.ServeContent(w, r, file, fi.ModTime(), f)
}
func (hv *HarmowareVisGo) getAssetsDir() http.Dir{
	currentRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	d := filepath.Join(currentRoot, "build")

	assetsDir := http.Dir(d)
	return assetsDir
}