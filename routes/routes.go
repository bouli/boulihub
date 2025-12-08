package routes

import (
	"fmt"
	"net/http"
)

type destiny struct {
	DestinyType string
	DestinyUrl  string
}

type website_host string
type website_path string
type websites map[website_host]map[website_path]destiny

func GetDestiny(r *http.Request, forceWebSite string) destiny {
	if forceWebSite != "" {
		r.Host = forceWebSite
	}
	fmt.Println("host: " + website_host(r.Host))
	fmt.Println("path: " + website_path(r.URL.Path))
	return getDestinies()[website_host(r.Host)][website_path(r.URL.Path)]
}

func GetFileDestiny(r *http.Request, forceWebSite string) destiny {
	if forceWebSite != "" {
		r.Host = forceWebSite
	}
	fmt.Println("host: " + website_host(r.Host))
	fmt.Println("pattern: " + website_path(r.Pattern))

	return getDestinies()[website_host(r.Host)][website_path(r.Pattern)]
}

func getDestinies() websites {
	myMap := make(websites)
	myMap["salviasupernova.com.br"] = make(map[website_path]destiny)
	myMap["salviasupernova.com.br"]["/"] = destiny{"template", "salviasupernova.com.br"}

	myMap["cesarbouli.com"] = make(map[website_path]destiny)
	myMap["cesarbouli.com"]["/"] = destiny{"template", "cesarbouli.com"}
	myMap["cesarbouli.com"]["/como-se-fosse-a-ultima-vez"] = destiny{"redirect", "https://www.youtube.com/watch?v=dQZYDKcQiEU"}
	myMap["cesarbouli.com"]["/cesar-bouli"] = destiny{"redirect", "https://distrokid.com/hyperfollow/cesarbouli/cesar-bouli"}
	myMap["cesarbouli.com"]["/hey-irmao"] = destiny{"redirect", "https://distrokid.com/hyperfollow/cesarbouli/hey-irmo-2"}
	myMap["cesarbouli.com"]["/fios-naturais"] = destiny{"redirect", "https://distrokid.com/hyperfollow/cesarbouli/fios-naturais"}
	myMap["cesarbouli.com"]["/vale-a-pena"] = destiny{"redirect", "https://distrokid.com/hyperfollow/cesarbouli/vale-a-pena-4"}

	myMap["byeceebee.com"] = make(map[website_path]destiny)
	myMap["byeceebee.com"]["/"] = destiny{"template", "byeceebee.com"}
	myMap["byeceebee.com"]["/diana"] = destiny{"redirect", "https://open.spotify.com/track/7GPSPivgaqYPQ6FstxXi9C"}
	myMap["byeceebee.com"]["/happy-end"] = destiny{"redirect", "https://open.spotify.com/album/427He8WE8uuO4ECdchYhh8"}
	myMap["byeceebee.com"]["/padroeira"] = destiny{"redirect", "https://open.spotify.com/album/3nSxhCNm1SY1XzAvblzqH1"}

	myMap["paralelossa.com.br"] = make(map[website_path]destiny)
	myMap["paralelossa.com.br"]["/"] = destiny{"template", "paralelossa.com.br"}
	myMap["paralelossa.com.br"]["/universo-surdo-e-mudo"] = destiny{"redirect", "https://open.spotify.com/album/30CfdUPREWbUZUk0TP9GVi"}

	myMap["bouli.com.br"] = make(map[website_path]destiny)
	myMap["bouli.com.br"]["/"] = destiny{"redirect", "https://cesarcardoso.cc"}

	myMap["littletalks.org"] = make(map[website_path]destiny)
	myMap["littletalks.org"]["/"] = destiny{"redirect", "https://github.com/LittleTalksOrg"}

	myMap["ohsanaworks.com"] = make(map[website_path]destiny)
	myMap["ohsanaworks.com"]["/"] = destiny{"static", "ohsanaworks.com/index.html"}
	myMap["ohsanaworks.com"]["/css/"] = destiny{"files", "ohsanaworks.com"}
	myMap["ohsanaworks.com"]["/download/"] = destiny{"files", "ohsanaworks.com"}
	myMap["ohsanaworks.com"]["/fonts/"] = destiny{"files", "ohsanaworks.com"}
	myMap["ohsanaworks.com"]["/images/"] = destiny{"files", "ohsanaworks.com"}
	myMap["ohsanaworks.com"]["/js/"] = destiny{"files", "ohsanaworks.com"}

	myMap["cesarcardoso.cc"] = make(map[website_path]destiny)
	myMap["cesarcardoso.cc"]["/"] = destiny{"static", "cesarcardoso.cc/index.html"}
	myMap["cesarcardoso.cc"]["/lebenslauf"] = destiny{"static", "cesarcardoso.cc/resume/resume-for-web-de.html"}
	myMap["cesarcardoso.cc"]["/cv"] = destiny{"static", "cesarcardoso.cc/resume/resume-for-web-en.html"}
	myMap["cesarcardoso.cc"]["/qr"] = destiny{"static", "cesarcardoso.cc/qr.html"}
	myMap["cesarcardoso.cc"]["/assets/"] = destiny{"files", "cesarcardoso.cc"}

	myMap["cesarcardoso.cc"]["/resume"] = destiny{"redirect", "/cv"}
	myMap["cesarcardoso.cc"]["/resume/en"] = destiny{"redirect", "/cv"}
	myMap["cesarcardoso.cc"]["/en/cv"] = destiny{"redirect", "/cv"}
	myMap["cesarcardoso.cc"]["/de/cv"] = destiny{"redirect", "/lebenslauf"}
	myMap["cesarcardoso.cc"]["/resume/de"] = destiny{"redirect", "/lebenslauf"}

	return myMap
}
