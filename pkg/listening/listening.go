package listening

import (
	"fmt"
	"log"
	"net/http"
)

func ListePrintServer(hs *http.Server) {
	fmt.Print(`
─────▄───▄
─▄█▄─█▀█▀█─▄█▄
▀▀████▄█▄████▀▀
─────▀█▀█▀
 	  __                                        
	 / _|                                       
 	| |___   _______  ___ __  _ __ ___  ___ ___ 
 	|  _\ \ / / _ \ \/ / '_ \| '__/ _ \/ __/ __|
 	| |  \ V /  __/>  <| |_) | | |  __/\__ \__ \
 	|_|   \_/ \___/_/\_\ .__/|_|  \___||___/___/
	                   | |                      
       		           |_|                      

`)

	log.Println("Listening..." + hs.Addr)
}
