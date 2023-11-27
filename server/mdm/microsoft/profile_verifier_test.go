package microsoft_mdm

import (
	"context"
	"testing"

	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/mdm/microsoft/syncml"
	"github.com/fleetdm/fleet/v4/server/mock"
	"github.com/stretchr/testify/require"
)

func TestLoopHostMDMLocURIs(t *testing.T) {
	ds := new(mock.Store)
	ctx := context.Background()

	ds.GetHostMDMProfilesExpectedForVerificationFunc = func(ctx context.Context, host *fleet.Host) (map[string]*fleet.ExpectedMDMProfile, error) {
		return map[string]*fleet.ExpectedMDMProfile{
			"N1": {Name: "N1", RawProfile: syncml.ForTestWithData(map[string]string{"L1": "D1"})},
			"N2": {Name: "N2", RawProfile: syncml.ForTestWithData(map[string]string{"L2": "D2"})},
			"N3": {Name: "N3", RawProfile: syncml.ForTestWithData(map[string]string{"L3": "D3", "L3.1": "D3.1"})},
		}, nil
	}

	type wantStruct struct {
		locURI      string
		data        string
		profileUUID string
		uniqueHash  string
	}
	got := []wantStruct{}
	err := LoopHostMDMLocURIs(ctx, ds, &fleet.Host{}, func(profile *fleet.ExpectedMDMProfile, hash, locURI, data string) {
		got = append(got, wantStruct{
			locURI:      locURI,
			data:        data,
			profileUUID: profile.Name,
			uniqueHash:  hash,
		})
	})
	require.NoError(t, err)
	require.ElementsMatch(
		t,
		[]wantStruct{
			{"L1", "D1", "N1", "1255198959"},
			{"L2", "D2", "N2", "2736786183"},
			{"L3", "D3", "N3", "894211447"},
			{"L3.1", "D3.1", "N3", "3410477854"},
		},
		got,
	)

}

func TestHashLocURI(t *testing.T) {
	testCases := []struct {
		name           string
		profileName    string
		locURI         string
		expectNotEmpty bool
	}{
		{
			name:           "basic functionality",
			profileName:    "profile1",
			locURI:         "uri1",
			expectNotEmpty: true,
		},
		{
			name:           "empty strings",
			profileName:    "",
			locURI:         "",
			expectNotEmpty: true,
		},
		{
			name:           "special characters",
			profileName:    "profile!@#",
			locURI:         "uri$%^",
			expectNotEmpty: true,
		},
		{
			name:           "long string input",
			profileName:    string(make([]rune, 1000)),
			locURI:         string(make([]rune, 1000)),
			expectNotEmpty: true,
		},
		{
			name:           "non-ASCII characters",
			profileName:    "プロファイル",
			locURI:         "URI",
			expectNotEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash := HashLocURI(tc.profileName, tc.locURI)
			if tc.expectNotEmpty {
				require.NotEmpty(t, hash, "hash should not be empty")
			}
		})
	}
}
