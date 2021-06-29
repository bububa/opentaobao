package hotelhstdf

// GetHotelRoomStaticParam 
type GetHotelRoomStaticParam struct {

    // 第1页
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页取100条
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 字典类型
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
