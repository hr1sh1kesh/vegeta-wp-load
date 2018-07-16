package src

func GetAPIEndPoints() []string {

	apiEndpoints := []string{"/wp-json/wp/v2/posts", "/wp-json/wp/v2/users", "/wp-json/wp/v2/categories"}

	return apiEndpoints

}
