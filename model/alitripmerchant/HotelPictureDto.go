package alitripmerchant

// HotelPictureDto 
type HotelPictureDto struct {

    // 类型名称
    
    TypeName   string `json:"type_name,omitempty" xml:"type_name,omitempty"`
    

    // 类型编码
    
    TypeCode   string `json:"type_code,omitempty" xml:"type_code,omitempty"`
    

    // 图片集合
    
    PictureAddress   []string `json:"picture_address,omitempty" xml:"picture_address>string,omitempty"`
    

}
