package viapi

// TextRectangle 
type TextRectangle struct {
    // 文字区域左上角x坐标
    Left   int64 `json:"left,omitempty" xml:"left,omitempty"`
    // 文字区域角度，角度范围[0, 360]
    Angle   int64 `json:"angle,omitempty" xml:"angle,omitempty"`
    // 文字区域左上角y坐标
    Top   int64 `json:"top,omitempty" xml:"top,omitempty"`
    // 文字区域高度
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    // 文字区域宽度
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
}
