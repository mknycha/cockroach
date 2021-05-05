package geomfn

import (
	"github.com/cockroachdb/cockroach/pkg/geo"
	"github.com/cockroachdb/cockroach/pkg/geo/geopb"
	"github.com/cockroachdb/errors"
	"github.com/twpayne/go-geom"
)

// Split TODO
func Split(a geo.Geometry, b geo.Geometry) (geo.Geometry, error) {
	if a.SRID() != b.SRID() {
		return geo.Geometry{}, geo.NewMismatchingSRIDsError(a.SpatialObject(), b.SpatialObject())
	}
	return split(a, b)
}

func split(g geo.Geometry, blade geo.Geometry) (geo.Geometry, error) {
	switch g.ShapeType() {
	case geopb.ShapeType_LineString:
		gt, err := g.AsGeomT()
		if err != nil {
			return geo.Geometry{}, err
		}
		gtLine := gt.(*geom.LineString)
		gtBlade, err := blade.AsGeomT()
		if err != nil {
			return geo.Geometry{}, err
		}
		return splitLine(gtLine, gtBlade)
	case geopb.ShapeType_Polygon:
		return geo.Geometry{}, errors.New("not implemented")
	case geopb.ShapeType_MultiPolygon:
		return geo.Geometry{}, errors.New("not implemented")
	case geopb.ShapeType_MultiLineString:
		return geo.Geometry{}, errors.New("not implemented")
	case geopb.ShapeType_GeometryCollection:
		return geo.Geometry{}, errors.New("not implemented")
		// TODO: Check what should happen for multipoint.
	default:
		return geo.Geometry{}, errors.Newf("input geometry type %s is not supported", g.ShapeType().String())
	}
}

func splitLine(gtLine *geom.LineString, gtBlade geom.T) (geo.Geometry, error) {
	switch gtTypedBlade := gtBlade.(type) {
	// TODO: Implement other types
	case *geom.Point:
		return splitLineByPointForSplit(gtLine, gtTypedBlade.Coords())
	default:
		// TODO: Can we find out what the geometry type is here, without adapting to Geometry?
		blade, err := geo.MakeGeometryFromGeomT(gtBlade)
		if err != nil {
			return geo.Geometry{}, errors.Newf("error parsing blade into Geometry: %w", err)
		}
		return geo.Geometry{}, errors.Newf("splitting a Line by %s is not supported", blade.ShapeType().String())
	}
}

func splitLineByPointForSplit(l *geom.LineString, p geom.Coord) (geo.Geometry, error) {
	_, lineStrings, err := splitLineByPoint(l, p)
	if err != nil {
		return geo.Geometry{}, err
	}
	multiLineString := geom.NewMultiLineString(l.Layout())
	for _, ls := range lineStrings {
		err := multiLineString.Push(ls)
		if err != nil {
			return geo.Geometry{}, err
		}
	}
	out, err := geo.MakeGeometryFromGeomT(multiLineString)
	if err != nil {
		return geo.Geometry{}, errors.Newf("could not transform output into geometry: %v", err)
	}
	return out, nil
}
