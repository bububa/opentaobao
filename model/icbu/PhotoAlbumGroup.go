package icbu

// PhotoAlbumGroup 结构体
type PhotoAlbumGroup struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// level1
	Level1 int64 `json:"level1,omitempty" xml:"level1,omitempty"`
	// level2
	Level2 int64 `json:"level2,omitempty" xml:"level2,omitempty"`
	// level3
	Level3 int64 `json:"level3,omitempty" xml:"level3,omitempty"`
}
