package alihouse

// ProjectReviewDraftDto 
type ProjectReviewDraftDto struct {

    // 楼盘id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 外部测评id
    
    OuterReviewId   string `json:"outer_review_id,omitempty" xml:"outer_review_id,omitempty"`
    

    // 菜鸟城市id
    
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    

    // 区域描述
    
    DistrictDesc   string `json:"district_desc,omitempty" xml:"district_desc,omitempty"`
    

    // 板块描述
    
    BlockDesc   string `json:"block_desc,omitempty" xml:"block_desc,omitempty"`
    

    // 轨道交通
    
    SubwayDesc   string `json:"subway_desc,omitempty" xml:"subway_desc,omitempty"`
    

    // 自驾道路
    
    MainRoad   string `json:"main_road,omitempty" xml:"main_road,omitempty"`
    

    // 教育配套
    
    EduSource   string `json:"edu_source,omitempty" xml:"edu_source,omitempty"`
    

    // 医疗配套
    
    MedicalSource   string `json:"medical_source,omitempty" xml:"medical_source,omitempty"`
    

    // 其他配套
    
    OtherSource   string `json:"other_source,omitempty" xml:"other_source,omitempty"`
    

    // 价格现状
    
    NowPrice   string `json:"now_price,omitempty" xml:"now_price,omitempty"`
    

    // 价格潜力
    
    PotentialPrice   string `json:"potential_price,omitempty" xml:"potential_price,omitempty"`
    

    // 户型分析
    
    HouseTypeAnalysis   string `json:"house_type_analysis,omitempty" xml:"house_type_analysis,omitempty"`
    

    // 样板房细节
    
    ModelHouse   string `json:"model_house,omitempty" xml:"model_house,omitempty"`
    

    // 公共交通
    
    PublicTraffic   string `json:"public_traffic,omitempty" xml:"public_traffic,omitempty"`
    

    // 亮点
    
    Lights   string `json:"lights,omitempty" xml:"lights,omitempty"`
    

    // 项目不足
    
    Defect   string `json:"defect,omitempty" xml:"defect,omitempty"`
    

    // 状态 1有效0 无效
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 发布时间 时间格式 yyyy-MM-dd HH:mm:ss
    
    PublishTime   string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
    

    // 1测试数据 0正常数据
    
    IsTest   int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
    

}
