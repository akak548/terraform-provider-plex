package plex

import (
	plexclient "github.com/akak548/go-plex-client"
	"testing"
)

var testuser1 = plexclient.Friends{
	ID:                        23333815,
	Title:                     "testuser1@hotmail.com",
	Thumb:                     "https://plex.tv/users/622cf3c172136286/avatar?c=1557445048",
	Protected:                 "0",
	Home:                      "0",
	AllowSync:                 "0",
	AllowCameraUpload:         "0",
	AllowChannels:             "0",
	FilterMovies:              "contentRating=G%2CTV-PG%2CTV-14%2CR%2CPG-13%2CPG%2CNC-17",
	Restricted:                "0",
	Username:                  "testuser1@hotmail.com",
	Email:                     "testuser1@hotmail.com",
	RecommendationsPlaylistID: "af7f2b02da76f7d2",
}

var testuser2 = plexclient.Friends{
	ID:                        23333816,
	Title:                     "testuser2@gmail.com",
	Thumb:                     "https://plex.tv/users/622cf3c172136286/avatar?c=1557445048",
	Protected:                 "0",
	Home:                      "0",
	AllowSync:                 "0",
	AllowCameraUpload:         "0",
	AllowChannels:             "0",
	FilterMovies:              "contentRating=G%2CTV-PG%2CTV-14%2CR%2CPG-13%2CPG%2CNC-17",
	Restricted:                "0",
	Username:                  "testuser2",
	Email:                     "testuser2@gmail.com",
	RecommendationsPlaylistID: "af7f2b02da76f7d2",
}

var testuser3 = plexclient.Friends{
	ID:                        23333817,
	Title:                     "testuser3@gmail.com",
	Thumb:                     "https://plex.tv/users/622cf3c172136286/avatar?c=1557445048",
	Protected:                 "0",
	Home:                      "0",
	AllowSync:                 "0",
	AllowCameraUpload:         "0",
	AllowChannels:             "0",
	FilterMovies:              "contentRating=G%2CTV-PG%2CTV-14%2CR%2CPG-13%2CPG%2CNC-17",
	Restricted:                "0",
	Username:                  "testuser3",
	Email:                     "testuser3@gmail.com",
	RecommendationsPlaylistID: "af7f2b02da76f7d2",
}

var friends = []plexclient.Friends{
	testuser1, testuser2, testuser3,
}

func TestFindFriend(t *testing.T) {
	var tests = []struct {
		input    string
		friends  []plexclient.Friends
		expected plexclient.Friends
		err      error
	}{
		// Test finding by username
		{
			input:    "testuser1@hotmail.com",
			friends:  friends,
			expected: testuser1,
			err:      nil,
		},
		// Test finding by email
		{
			input:    "testuser2@gmail.com",
			friends:  friends,
			expected: testuser2,
			err:      nil,
		},
		// Test friendNotFound error
		{
			input:    "notfound@gmail.com",
			friends:  friends,
			expected: plexclient.Friends{},
			err:      friendNotFound{"notfound@gmail.com"},
		},
	}

	for _, test := range tests {
		i, err := findFriend(test.input, friends)

		if i != test.expected {
			t.Errorf("Output Received %v, and expected %v", i, test.expected)
		}
		if err != test.err {
			t.Errorf("Received %v, and expected %v for error", err, test.err)
		}
	}
}
