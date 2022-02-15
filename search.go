package groupietracker

import (
	"fmt"
	"os"
	"strconv"
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
		fmt.Println(err)
		os.Exit(2)
	}
	for _, artist := range data {
		if intSearchCreationYear == artist.CreationDate {
			artistAllData = append(artistAllData, artist)
		}
		return artistAllData
	}
	return nil
}

func SearchByFirstAlbum(data []ArtistAllData, searchFirstAlbum string) []ArtistAllData {
	var artistAllData []ArtistAllData
	for _, artist := range data {
		if searchFirstAlbum == artist.FirstAlbum {
			artistAllData = append(artistAllData, artist)
		}
		return artistAllData
	}
	return nil
}

func SearchByName(data []ArtistAllData, searchName string) []ArtistAllData {
	var artistAllData []ArtistAllData
	for _, artist := range data {
		if searchName == artist.Name {
			artistAllData = append(artistAllData, artist)
		}
		return artistAllData
	}
	return nil
}
