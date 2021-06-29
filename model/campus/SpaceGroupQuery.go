package campus

// SpaceGroupQuery 
type SpaceGroupQuery struct {

    // 分页限制
    
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    

    // 分组ID集合
    
    Ids   []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
    

    // 分组ID
    
    GroupId   int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
    

    // 楼层ID
    
    FloorId   int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
    

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 公司ID
    
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    

    // 类别ID
    
    TypeId   int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
    

    // 园区ID
    
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    

    // 楼宇ID
    
    BuildingId   int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
    

    // 类型编码
    
    TypeCode   string `json:"type_code,omitempty" xml:"type_code,omitempty"`
    

    // 空间分组编码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 分组名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 分组名称或者code
    
    NameOrCode   string `json:"name_or_code,omitempty" xml:"name_or_code,omitempty"`
    

}
