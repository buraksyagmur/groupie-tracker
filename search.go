package groupietracker

import (
	"fmt"
	"strconv"
	"strings"
)

// func SearchFull(data []ArtistAllData, SearchCreation int, SearchFirst, SearchName string) []ArtistAllData {
// 	var artistAllData []ArtistAllData
// 	for _, artist := range data {
// 		if SearchByCreationYear(SearchCreation, artist.CreationDate) &&
// 			SearchByFirstAlbum(SearchFirst, artist.FirstAlbum) &&
// 			SearchByName(SearchName, artist.Name) {
// 			artistAllData = append(artistAllData, artist)
// 		}
// 	}
// 	return artistAllData
// }

func SearchByCreationYear(data []ArtistAllData, searchCreationYear string) []ArtistAllData {
	var artistAllData []ArtistAllData
	intSearchCreationYear, err := strconv.Atoi(searchCreationYear)
	if err != nil {
		return nil
	}
	for _, artist := range data {
		if intSearchCreationYear == artist.CreationDate {
			artistAllData = append(artistAllData, artist)
		}
	}
	fmt.Println("searchedbycreationyear")
	return artistAllData
}

func SearchByFirstAlbum(data []ArtistAllData, searchFirstAlbum string) []ArtistAllData {
	var artistAllData []ArtistAllData
	for _, artist := range data {
		if searchFirstAlbum == artist.FirstAlbum {
			artistAllData = append(artistAllData, artist)
		}
	}
	fmt.Println("searchedbyfirstalbum")
	return artistAllData
}

func SearchByName(data []ArtistAllData, searchName string) []ArtistAllData {
	var artistAllData []ArtistAllData
	for _, artist := range data {
		if searchName == artist.Name {
			artistAllData = append(artistAllData, artist)
		}

		

	}
	fmt.Println(artistAllData, "------")
	if artistAllData == nil {
		capname := strings.Title(searchName)
		for _, artist := range data {
			if capname == artist.Name {
				artistAllData = append(artistAllData, artist)
				fmt.Println("itsduplicate")
			}
		}
	}
	fmt.Println("searchedbyname")
	return artistAllData
}

func SearchByLocation(data []ArtistAllData, searchLocation string) []ArtistAllData {
	var artistAllData []ArtistAllData
	if searchLocation == "queen" || searchLocation == "Queen" {
		searchLocation = "queensland-australia"
	}
	for _, artist := range data {
		for i := 0; i < len(artist.Locations); i++ {
			if searchLocation == artist.Locations[i] {
				artistAllData = append(artistAllData, artist)
			}
		}
	}
	fmt.Println(searchLocation)
	fmt.Println("searchedbylocations")
	return artistAllData
}

func SearchByMember(data []ArtistAllData, searchMember string) []ArtistAllData {
	var artistAllData []ArtistAllData
	for _, artist := range data {
		for i := 0; i < len(artist.Members); i++ {
			if searchMember == artist.Members[i] {
				artistAllData = append(artistAllData, artist)
			}
		}
	}
	if artistAllData == nil {
		capmember := strings.Title(searchMember)
		for _, artist := range data {
			for i := 0; i < len(artist.Members); i++ {
				if capmember == artist.Members[i] {
					artistAllData = append(artistAllData, artist)
				}
			}
		}
	}

	fmt.Println("searchedbymembers")
	return artistAllData
}
