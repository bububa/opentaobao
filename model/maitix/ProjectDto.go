package maitix

// ProjectDto 
type ProjectDto struct {

    // 城市信息
    
    City   *IdNameDto `json:"city,omitempty" xml:"city,omitempty"`
    

    // 项目大类编码-废弃,使用项目详情结果的内容
    
    ClassifyCode   string `json:"classify_code,omitempty" xml:"classify_code,omitempty"`
    

    // 项目大类名称-废弃,使用项目详情结果的内容
    
    ClassifyName   string `json:"classify_name,omitempty" xml:"classify_name,omitempty"`
    

    // 国家
    
    Country   *IdNameDto `json:"country,omitempty" xml:"country,omitempty"`
    

    // 项目介绍，目前该字段废弃
    
    Introduce   string `json:"introduce,omitempty" xml:"introduce,omitempty"`
    

    // 是否有座： 0=无座 1=有座
    
    IsHasSeat   int64 `json:"is_has_seat,omitempty" xml:"is_has_seat,omitempty"`
    

    // 项目海报图片地址，目前该字段废弃
    
    PosterUrl   string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
    

    // 项目ID
    
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    

    // 项目名称
    
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    

    // 项目状态，0：创建中 10：已创建 20：待销售 30：销售中 40：项目取消 50：项目结束 60:定时开售，一般30之前的状态不会透出给渠道
    
    ProjectStatus   int64 `json:"project_status,omitempty" xml:"project_status,omitempty"`
    

    // 项目类型 0:普通项目-废弃
    
    ProjectType   int64 `json:"project_type,omitempty" xml:"project_type,omitempty"`
    

    // 省
    
    Province   *IdNameDto `json:"province,omitempty" xml:"province,omitempty"`
    

    // 备注-废弃
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 项目子类编码-废弃,使用项目详情结果的内容
    
    SubClassifyCode   string `json:"sub_classify_code,omitempty" xml:"sub_classify_code,omitempty"`
    

    // 项目子类名称-废弃,使用项目详情结果的内容
    
    SubClassifyName   string `json:"sub_classify_name,omitempty" xml:"sub_classify_name,omitempty"`
    

    // 项目三级分类编码-废弃,使用项目详情结果的内容
    
    ThirdClassifyCode   string `json:"third_classify_code,omitempty" xml:"third_classify_code,omitempty"`
    

    // 项目三级分类名称-废弃,使用项目详情结果的内容
    
    ThirdClassifyName   string `json:"third_classify_name,omitempty" xml:"third_classify_name,omitempty"`
    

    // 主办方ID
    
    TraderIdsArrList   []int64 `json:"trader_ids_arr_list,omitempty" xml:"trader_ids_arr_list>int64,omitempty"`
    

    // 主办方名称
    
    TraderNamesArrList   []string `json:"trader_names_arr_list,omitempty" xml:"trader_names_arr_list>string,omitempty"`
    

    // 场馆
    
    Venue   *VenueDto `json:"venue,omitempty" xml:"venue,omitempty"`
    

    // 是否测试项目 0-正式项目 1-测试项目
    
    IsTest   int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
    

}
