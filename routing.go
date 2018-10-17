package tomtom

import (
	"fmt"
)

type RoutingService struct {
	client *Client
}

type CalculateRouteResponse struct {
	Copyright     string
	FormatVersion string
	Privacy       string
	Routes        []Route
}

type Route struct {
	Summary  Summary
	Legs     []Leg
	Sections []Section
}

type Leg struct {
	Summary Summary
	Points  []Point
}

type Summary struct {
	LengthInMeters        uint
	TravelTimeInSeconds   uint
	TrafficDelayInSeconds uint
	DepartureTime         string
}

type Point struct {
	Latitude  float32
	Longitude float32
}

type Section struct {
	StartPointIndex uint
	EndPointIndex   uint
	SectionType     string
	TravelMode      string
}

func (s *RoutingService) CalculateRoute(from, to Point) (*CalculateRouteResponse, error) {
	locations := fmt.Sprintf("%f,%f:%f,%f", from.Latitude, from.Longitude, to.Latitude, to.Longitude)
	url := fmt.Sprintf(
		"/routing/%d/%s/%s/%s",
		s.client.ApiVersion,
		"calculateRoute",
		locations,
		s.client.ContentType,
	)
	req, err := s.client.newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	var res CalculateRouteResponse
	_, err = s.client.do(req, &res)
	return &res, err
}
