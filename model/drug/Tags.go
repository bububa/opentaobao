package drug

// Tags 
type Tags struct {
    // type
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // picPath
    PicPath   string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
    // desc
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // shortDesc
    ShortDesc   string `json:"short_desc,omitempty" xml:"short_desc,omitempty"`
    // manFanType
    ManFanType   string `json:"man_fan_type,omitempty" xml:"man_fan_type,omitempty"`
}
