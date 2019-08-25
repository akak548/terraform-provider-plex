provider "plex" {
	api_address = "http://192.168.2.68:32400"
}

resource "plex_friend" "TestUser" {
	username = "akak548test"
	email_address = "akak548+plex@gmail.com"
	machine_id = "78906236672345bacd073d086b4240d02e689b40" 
	allow_cameraupload = "0"
	allow_channels = "0"
	allow_sync = "0"
}
