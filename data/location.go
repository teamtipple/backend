package data 

import (
    "time"
    "github.com/dancannon/gorethink"
)

type Location struct {
    Name string `gorethink:"name"`
    Location gorethink.Point `gorethink:"location"`
    Expires time.Duration
}

func GetLocations(lat float64, long float64, rad float64) ([]Location, error) {
    return nil, nil
}

func PushLocations([]Location) error {
    return nil    
}
