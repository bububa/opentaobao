package alihouse

// SyncProjectLayoutDto 
type SyncProjectLayoutDto struct {

    // 户型图片多张以英文逗号隔开
    
    LayoutImages   string `json:"layout_images,omitempty" xml:"layout_images,omitempty"`
    

    // 户型均价
    
    AvgPrice   string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
    

    // 户型总价
    
    TotalPrice   string `json:"total_price,omitempty" xml:"total_price,omitempty"`
    

    // 室内面积
    
    InsideArea   string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
    

    // 建筑面积
    
    ConstructionArea   string `json:"construction_area,omitempty" xml:"construction_area,omitempty"`
    

    // 卫数量
    
    Bathroom   string `json:"bathroom,omitempty" xml:"bathroom,omitempty"`
    

    // 厅数量
    
    Hall   string `json:"hall,omitempty" xml:"hall,omitempty"`
    

    // 室数量
    
    Room   string `json:"room,omitempty" xml:"room,omitempty"`
    

    // 楼盘ID
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 户型外部唯一ID
    
    OuterLayoutId   string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
    

    // 户型名称
    
    LayoutName   string `json:"layout_name,omitempty" xml:"layout_name,omitempty"`
    

    // 朝向
    
    Orientation   int64 `json:"orientation,omitempty" xml:"orientation,omitempty"`
    

    // 描述说明
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 是否删除 1 是 0 否
    
    IsDeleted   string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
    

}
