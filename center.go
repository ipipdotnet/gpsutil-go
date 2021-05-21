package gpsutil

import (
	"fmt"
	"github.com/golang/geo/s2"
)

// 22.550060,113.937950|22.550200,113.938230
func Center()  {

	p1 := s2.PointFromLatLng(s2.LatLngFromDegrees(22.550200,113.938230))
	p2 := s2.PointFromLatLng(s2.LatLngFromDegrees(22.550060,113.937950))
	p3 := s2.PointFromLatLng(s2.LatLngFromDegrees(22.543270,113.930480))

	pts := s2.LoopFromPoints([]s2.Point{
		p1, p2, p3,
	})
	fmt.Println(pts.Centroid().Distance(s2.PointFromLatLng(s2.LatLngFromDegrees(22.550200,113.938230))))
	fmt.Println(pts.Centroid().Distance(s2.PointFromLatLng(s2.LatLngFromDegrees(22.550060,113.937950))))
	fmt.Println(pts.Centroid().Distance(s2.PointFromLatLng(s2.LatLngFromDegrees(22.550200,113.938060))))
	fmt.Println(s2.LatLngFromPoint(pts.Centroid()))

}