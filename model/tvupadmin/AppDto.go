package tvupadmin

// AppDto 
type AppDto struct {
    // 应用ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // packageName
    PackageName   string `json:"package_name,omitempty" xml:"package_name,omitempty"`
}
