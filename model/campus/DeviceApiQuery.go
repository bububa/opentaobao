package campus

// DeviceApiQuery 
type DeviceApiQuery struct {

    // 分页大小
    
    Limit   int64 `json:"limit,omitempty" xml:"limit,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 园区id
    
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    

    // 模糊查询关键字
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
    

    // 设备模板code
    
    TemplateCode   string `json:"template_code,omitempty" xml:"template_code,omitempty"`
    

    // 园区id集合
    
    CampusIdList   []int64 `json:"campus_id_list,omitempty" xml:"campus_id_list>int64,omitempty"`
    

    // 模板code集合,具体设备模板code信息请查阅‘平台技术’下‘设备详细信息开发文档’。[生成参考文档该字段入参是数组，有误。入参类型请使用java.lang.List类型]
    
    TemplateCodeList   []string `json:"template_code_list,omitempty" xml:"template_code_list>string,omitempty"`
    

    // 楼层id
    
    FloorId   int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
    

    // 空间单元id
    
    SpaceId   int64 `json:"space_id,omitempty" xml:"space_id,omitempty"`
    

    // 设备uuid集合[生成参考文档该字段入参是数组，有误。入参类型请使用java.lang.List类型]
    
    UuidList   []string `json:"uuid_list,omitempty" xml:"uuid_list>string,omitempty"`
    

    // 楼宇id
    
    BuildingId   int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
    

    // 空间单元id集合[生成参考文档该字段入参是数组，有误。入参类型请使用java.lang.List类型]
    
    SpaceIdList   []int64 `json:"space_id_list,omitempty" xml:"space_id_list>int64,omitempty"`
    

    // 是否启用
    
    BeRun   bool `json:"be_run,omitempty" xml:"be_run,omitempty"`
    

    // 0->在线 1->离线 2->故障
    
    RunStatus   int64 `json:"run_status,omitempty" xml:"run_status,omitempty"`
    

    // 根据标签名称精确查询
    
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    

    // 模糊查询设备code或名称
    
    NameOrCode   string `json:"name_or_code,omitempty" xml:"name_or_code,omitempty"`
    

    // 根据设备code精确匹配
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 楼层id集合
    
    FloorIdList   []int64 `json:"floor_id_list,omitempty" xml:"floor_id_list>int64,omitempty"`
    

    // 楼宇id集合
    
    BuildingIdList   []int64 `json:"building_id_list,omitempty" xml:"building_id_list>int64,omitempty"`
    

}
