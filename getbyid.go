package groupietracker

import "errors"

func GetArtistByID(id int) (ArtistDataWithRelation, error) {
	for _, artist := range Artists {
		if artist.ID == id {
			return artist, nil
		}
	}
	return ArtistDataWithRelation{}, errors.New("artist ID not found")
}

func GetDateByID(id int) (DateAllData, error) {
	for _, date := range Dates.Index {
		if date.ID == id {
			return date, nil
		}
	}
	return DateAllData{}, errors.New("date ID not found")
}

func GetLocationByID(id int) (LocationAllData, error) {
	for _, location := range Locations.Index {
		if location.ID == id {
			return location, nil
		}
	}
	return LocationAllData{}, errors.New("location ID not found")
}

func GetRelationByID(id int) (RelationAllData, error) {
	for _, relation := range Relations.Index {
		if relation.ID == id {
			return relation, nil
		}
	}
	return RelationAllData{}, errors.New("relation ID not found")
}

func GetFullDataById(id int) (ArtistAllData, error) {
	for _, artist := range ArtistsFull {
		if artist.ID == id {
			return artist, nil
		}
	}
	return ArtistAllData{}, errors.New("artist full data ID not found")
}
