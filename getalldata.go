package groupietracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetArtistsData() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return errors.New("error - get artist")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error- read artist")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

func GetDatesData() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return errors.New("error - get dates")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error - read dates")
	}
	json.Unmarshal(bytes, &Dates)
	return nil
}

func GetLocationsData() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return errors.New("error - get location")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error- read location")
	}
	json.Unmarshal(bytes, &Locations)
	return nil
}

func GetRelationsData() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return errors.New("error - get relation")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error - read relation")
	}
	json.Unmarshal(bytes, &Relations)
	return nil
}

func GetData() error {
	if GetArtistsData() != nil || GetLocationsData() != nil || GetDatesData() != nil || GetRelationsData() != nil {
		return errors.New("error - get artist or location or date or relation")
	}
	if ArtistsFull != nil {
		fmt.Println("duplication data")
	} else {
		for i := range Artists {
			var getdata ArtistAllData
			getdata.ID = i + 1
			getdata.Image = Artists[i].Image
			getdata.Name = Artists[i].Name
			getdata.Members = Artists[i].Members
			getdata.CreationDate = Artists[i].CreationDate
			getdata.FirstAlbum = Artists[i].FirstAlbum
			getdata.Locations = Locations.Index[i].Locations
			getdata.ConcertDates = Dates.Index[i].Dates
			getdata.DatesLocations = Relations.Index[i].DatesLocations
			ArtistsFull = append(ArtistsFull, getdata)

		}
	}
	return nil
}
