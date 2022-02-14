package groupietracker

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// type filterItems struct {
// 	StartYear         int
// 	FinishYear        int
// 	FirstAlbTime      time.Time
// 	FirstAlbUntilTime time.Time
// 	NumMembers        int
// }

func FilterCreation(data []ArtistAllData, start, end, first, firstuntil, mem, loc string) []ArtistAllData {
	// var filterItemsData filterItems

	if start == "" || end == "" || mem == "" {
		start = "1900"
		end = "2022"

	}
	startyear, err1 := strconv.Atoi(start)
	finishyear, err2 := strconv.Atoi(end)
	// members, err3 := strconv.Atoi(mem)

	if err1 != nil || err2 != nil {
		log.Fatal("error - check the years")
	}
	// filterItemsData.StartYear = startyear
	// filterItemsData.FinishYear = finishyear
	// filterItemsData.NumMembers = members

	firstInTime, _ := time.Parse("2006-01-02", first)
	firstUntilInTime, _ := time.Parse("2006-01-02", firstuntil)

	// filterItemsData.FirstAlbTime = firstInTime
	// filterItemsData.FirstAlbUntilTime = firstUntilInTime

	fmt.Println(firstInTime)
	fmt.Println(firstUntilInTime)
	// fmt.Println(filterItemsData)

	var artistAllData []ArtistAllData

	for _, artist := range data {
		firstAlbInTime, _ := time.Parse("02-01-2006", artist.FirstAlbum)
		if filterByYear(artist.CreationDate, startyear, finishyear) &&
			filterByDate(firstAlbInTime, firstInTime, firstUntilInTime) &&
			filterByNumMem(mem, len(artist.Members)) &&
			filterByLocation(loc, artist.Locations) {
			artistAllData = append(artistAllData, artist)
		}
	}
	return artistAllData
}

func filterByYear(crtyear, start, end int) bool {
	if crtyear >= start && crtyear <= end {
		return true
	}
	return false
}

func filterByDate(firstAlbDate, first, firstUntil time.Time) bool {
	if firstAlbDate.After(first) && firstAlbDate.Before(firstUntil) {
		return true
	}
	return false
}

func filterByNumMem(filterNum string, numMem int) bool {
	var srune []rune
	filterNum = strings.Join(strings.Fields(filterNum), "")
	srune = []rune(filterNum)
	for i := 0; i < len(filterNum); i++ {
		intoffilter, err := strconv.Atoi(string(srune[i]))
		if err != nil {
			log.Fatal("error - member filter")
		}
		if intoffilter == numMem {
			return true
		}
	}

	return false
}

func filterByLocation(filterLoca string, location []string) bool {
	if filterLoca == "" {
		return true
	}
	for i := 0; i < len(location); i++ {
		if filterLoca == location[i] {
			return true
		}
	}
	return false
}
