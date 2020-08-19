package raytrace

type Image struct {
	Width int `json:"width"`
	Height int `json:"height"`
}

type Camera struct {
	Position Vector3 `json:"position"`
}

type Scene struct {
	Camera Camera `json:"camera"`
	Spheres []Sphere `json:"spheres"`
}

type Config struct {
	Image Image `json:"image"`
	Scene Scene `json:"scene"`
}
