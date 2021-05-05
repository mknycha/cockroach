package geomfn

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/geo"
	"github.com/stretchr/testify/require"
)

func TestSplit(t *testing.T) {
	t.Run("split line by point", func(t *testing.T) {
		got, err := Split(
			geo.MustParseGeometry("LINESTRING(0 0,10 10)"),
			geo.MustParseGeometry("POINT(1.75 1.75)"),
		)
		require.NoError(t, err)
		want := geo.MustParseGeometry("MULTILINESTRING((0 0, 1.75 1.75), (1.75 1.75, 10 10))")
		require.Equal(t, want, got)
	})

	t.Run("Geometries with different SRIDs", func(t *testing.T) {
		_, err := Split(
			geo.MustParseGeometry("SRID=4269;LINESTRING(0 0,5 7.5)"),
			geo.MustParseGeometry("SRID=4326;LINESTRING(0 0,5 7.5)"),
		)
		require.Error(t, err)
		require.Equal(t, "operation on mixed SRIDs forbidden: (LineString, 4269) != (LineString, 4326)", err.Error())
	})

	t.Run("unsupported Geometry type", func(t *testing.T) {
		_, err := Split(
			geo.MustParseGeometry("MULTIPOINT(0 0, 100 100)"),
			geo.MustParseGeometry("LINESTRING(0 0,5 7.5)"),
		)
		require.Error(t, err)
		require.Equal(t, "input geometry type MultiPoint is not supported", err.Error())
	})
}
