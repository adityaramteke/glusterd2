package e2e

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gluster/glusterd2/pkg/api"
	georepapi "github.com/gluster/glusterd2/plugins/georeplication/api"
	"github.com/stretchr/testify/require"
)

func TestGeorepCreate(t *testing.T) {
	r := require.New(t)

	gds, err := setupCluster("./config/1.yaml", "./config/2.yaml")
	r.Nil(err)
	defer teardownCluster(gds)

	brickDir, err := ioutil.TempDir("", "TestGeorepCreate")
	r.Nil(err)
	defer os.RemoveAll(brickDir)

	var brickPaths []string
	for i := 1; i <= 4; i++ {
		brickPath, err := ioutil.TempDir(brickDir, "brick")
		r.Nil(err)
		brickPaths = append(brickPaths, brickPath)
	}

	client := initRestclient(gds[0].ClientAddress)

	volname1 := "testvol1"
	reqVol := api.VolCreateReq{
		Name: volname1,
		Bricks: []string{
			gds[0].PeerID() + ":" + brickPaths[0],
			gds[0].PeerID() + ":" + brickPaths[1]},
		Force: true,
	}
	vol1, err := client.VolumeCreate(reqVol)
	r.Nil(err)

	volname2 := "testvol2"
	reqVol = api.VolCreateReq{
		Name: volname2,
		Bricks: []string{
			gds[1].PeerID() + ":" + brickPaths[2],
			gds[1].PeerID() + ":" + brickPaths[3]},
		Force: true,
	}
	vol2, err := client.VolumeCreate(reqVol)
	r.Nil(err)

	reqGeorep := georepapi.GeorepCreateReq{
		MasterVol:  volname1,
		SlaveVol:   volname2,
		SlaveHosts: []string{gds[1].PeerAddress},
	}

	_, err = client.GeorepCreate(vol1.ID.String(), vol2.ID.String(), reqGeorep)
	r.Nil(err)

	// delete volume
	err = client.VolumeDelete(volname1)
	r.Nil(err)

	// delete volume
	err = client.VolumeDelete(volname2)
	r.Nil(err)
}
