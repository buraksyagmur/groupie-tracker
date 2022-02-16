package main

import (
	"fmt"
	"groupietracker"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var data []groupietracker.ArtistAllData

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "GET" {
		err := groupietracker.GetData()
		if err != nil {
			fmt.Println(1)
			log.Fatal("error - get data function")
		}
	}
	data = groupietracker.ArtistsFull
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Parse Error: %v", err)
		http.Error(w, "Error when Parsing", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, data); err != nil {
		log.Printf("Execute Error: %v", err)
		http.Error(w, "Error when Executing", http.StatusInternalServerError)
		return
	}
}

func detailspage(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("details")
	idInt, _ := strconv.Atoi(idString)
	artist, _ := groupietracker.GetFullDataById(idInt)

	tpl, err := template.ParseFiles("templates/fulldetails.html")
	if err != nil {
		fmt.Println(4)
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tpl.Execute(w, artist); err != nil {
		fmt.Println(5)
		http.Error(w, err.Error(), 400)
		return
	}
}

func checkNoDup(newcomer string, existing []string) bool {
	for i := 0; i < len(existing); i++ {
		if newcomer == existing[i] {
			return false
		}
	}
	return true
}

func filterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "GET" {
		locs := []string{}
		for numArtist := 0; numArtist < len(data)-1; numArtist++ {
			// fmt.Print(len(data[numArtist].Locations))
			for loc := 0; loc < len(data[numArtist].Locations)-1; loc++ {
				if checkNoDup(data[numArtist].Locations[loc], locs) {
					locs = append(locs, data[numArtist].Locations[loc])
				}
			}
		}
		// fmt.Print(locs)

		// get the filter form
		// use filepath.Join
		tpl, err := template.New("filtersGet.html").Funcs(template.FuncMap{
			"Rename": func(str string) string {
				newStr := strings.Replace(str, "-", ", ", -1)
				return strings.Replace(newStr, "_", " ", -1)
			},
		}).ParseFiles("templates/filtersGet.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if err := tpl.Execute(w, locs); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
	if r.Method == "POST" {
		// process the input
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request in Form", http.StatusBadRequest)
			return
		}
		filtercreate := r.FormValue("create")
		filtercreateuntil := r.FormValue("createuntil")
		filterfirst := r.FormValue("first")
		filterfirstuntil := r.FormValue("firstuntil")
		filtermembers1 := r.FormValue("member1")
		filtermembers2 := r.FormValue("member2")
		filtermembers3 := r.FormValue("member3")
		filtermembers4 := r.FormValue("member4")
		filtermembers5 := r.FormValue("member5")
		filtermembers6 := r.FormValue("member6")
		filtermembers7 := r.FormValue("member7")
		filtermembers8 := r.FormValue("member8")
		filtermembers9 := r.FormValue("member9")
		filterlocation := r.FormValue("locations")
		filterallmembers := (filtermembers1 + filtermembers2 + filtermembers3 + filtermembers4 + filtermembers5 + filtermembers6 + filtermembers7 + filtermembers8 + filtermembers9)
		filteredData := groupietracker.FilterCreation(data, filtercreate, filtercreateuntil, filterfirst, filterfirstuntil, filterallmembers, filterlocation)
		tpl, err := template.ParseFiles("templates/filtered.html")
		if err != nil {
			log.Printf("Parse Error: %v", err)
			http.Error(w, "Error when Parsing", http.StatusInternalServerError)
			return
		}

		if err := tpl.Execute(w, filteredData); err != nil {
			log.Printf("Execute Error: %v", err)
			http.Error(w, "Error when Executing", http.StatusInternalServerError)
			return
		}

	}
}

func searchPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// err := groupietracker.GetData()
		// if err != nil {
		// 	fmt.Println(1)
		// 	log.Fatal("error - get data function")
		// }
	}
	// searchwords:=r.FormValue("searchbar")
	// afterSearchData:=groupietracker.SearchFull(data, )
	var searchedData []groupietracker.ArtistAllData
	searchvalue := r.FormValue("searchbar")
	fmt.Println(searchvalue)
	searchvalueslice := strings.Split(searchvalue, "")
	searchvaluesrune := []rune(searchvalue)
	if searchvalueslice[1] == "-" || searchvalueslice[2] == "-" {
		fmt.Println("first album date")
		searchedData = groupietracker.SearchByFirstAlbum(data, searchvalue)
	} else if searchvalueslice[0] == "1" || searchvalueslice[0] == "2" {
		fmt.Println("this value is creation year")
		searchedData = groupietracker.SearchByCreationYear(data, searchvalue)
	} else if searchvaluesrune[0] >= 65 && searchvaluesrune[0] <= 90 {
		searchedData = groupietracker.SearchByName(data, searchvalue)
		fmt.Println("band name")
		if searchedData == nil || searchvalue == "Phil Collins" {
			searchedData = groupietracker.SearchByMember(data, searchvalue)
			fmt.Println("member")
		}
	} else if searchvaluesrune[0] >= 97 && searchvaluesrune[0] <= 122 {
		searchedData = groupietracker.SearchByLocation(data, searchvalue)
		fmt.Println("location name")
	}
	tpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		log.Printf("Parse Error: %v", err)
		http.Error(w, "Error when Parsing", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, searchedData); err != nil {
		log.Printf("Execute Error: %v", err)
		http.Error(w, "Error when Executing", http.StatusInternalServerError)
		return
	}
}

func main() {
	err := exec.Command("xdg-open", "http://localhost:8080/").Start()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("style"))
	fmt.Printf("Starting server at port 8080\n")

	mux := http.NewServeMux()
	mux.Handle("/mainfunction/style/", http.StripPrefix("/mainfunction/style/", fs))
	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/details", detailspage)
	mux.HandleFunc("/filters", filterPage)
	mux.HandleFunc("/search", searchPage)

	// add not found page
	// mux.HandleFunc("/filtered", filteredPage)
	http.ListenAndServe(":8080", mux)
}
