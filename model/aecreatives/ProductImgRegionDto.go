package aecreatives

// ProductImgRegionDto 
type ProductImgRegionDto struct {
    // 图片识别的坐标 pos_top_left_x
    PosTopLeftX   string `json:"pos_top_left_x,omitempty" xml:"pos_top_left_x,omitempty"`
    // 图片识别的坐标 pos_top_left_y
    PosTopLeftY   string `json:"pos_top_left_y,omitempty" xml:"pos_top_left_y,omitempty"`
    // 图片识别的坐标 pos_bottom_right_x
    PosBottomRightX   string `json:"pos_bottom_right_x,omitempty" xml:"pos_bottom_right_x,omitempty"`
    // 图片识别的坐标 pos_bottom_right_y
    PosBottomRightY   string `json:"pos_bottom_right_y,omitempty" xml:"pos_bottom_right_y,omitempty"`
}
