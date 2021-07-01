package feedflow

// ResultDTO 
type ResultDTO struct {
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 报表信息
    RptList   []RptResultDTO `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 总数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回信息
    ResultCode   *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 单元id
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
    // 广告位列表
    AdzoneBindList   []AdzoneBindDTO `json:"adzone_bind_list,omitempty" xml:"adzone_bind_list>adzone_bind_dto,omitempty"`
    // 绑定创意的列表
    CreativeBindList   []CreativeBindDTO `json:"creative_bind_list,omitempty" xml:"creative_bind_list>creative_bind_dto,omitempty"`
    // 返回数据结果
    Results   []AdgroupDTO `json:"results,omitempty" xml:"results>adgroup_dto,omitempty"`
    // 广告位列表
    AdzoneList   []AdzoneDTO `json:"adzone_list,omitempty" xml:"adzone_list>adzone_dto,omitempty"`
    // 人群列表
    Crowds   []CrowdDTO `json:"crowds,omitempty" xml:"crowds>crowd_dto,omitempty"`
    // 错误信息
    ErrorList   []ErrorInfoDTO `json:"error_list,omitempty" xml:"error_list>error_info_dto,omitempty"`
    // 商品列表
    ItemList   []ItemDTO `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
    // 标签信息
    Labels   []LabelDTO `json:"labels,omitempty" xml:"labels>label_dto,omitempty"`
    // 定向结构
    Targets   []TargetDTO `json:"targets,omitempty" xml:"targets>target_dto,omitempty"`
}
