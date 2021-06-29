package campus

// Floor 
type Floor struct {

    // orderNo
    
    OrderNo   int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
    

    // description
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // buildingName
    
    BuildingName   string `json:"building_name,omitempty" xml:"building_name,omitempty"`
    

    // buildingId
    
    BuildingId   int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
    

    // campusName
    
    CampusName   string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
    

    // campusId
    
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    

    // companyId
    
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    

    // 楼层编码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 楼层名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // modifier
    
    Modifier   string `json:"modifier,omitempty" xml:"modifier,omitempty"`
    

    // creator
    
    Creator   string `json:"creator,omitempty" xml:"creator,omitempty"`
    

    // gmtModified
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // gmtCreate
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // isDelete
    
    IsDelete   bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
    

    // 是否启用
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 地图楼层id
    
    GeoFloorId   int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
    

    // uuid
    
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    

    // 面积
    
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    

    // 高度
    
    Height   string `json:"height,omitempty" xml:"height,omitempty"`
    

    // 高度
    
    HeightStr   string `json:"height_str,omitempty" xml:"height_str,omitempty"`
    

    // 面积
    
    AreaStr   string `json:"area_str,omitempty" xml:"area_str,omitempty"`
    

}
