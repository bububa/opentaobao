package maitix

// ProjectInfoDTO 
type ProjectInfoDTO struct {
    // 演出城市
    City   *IdNameDTO `json:"city,omitempty" xml:"city,omitempty"`
    // 项目大类编码-取项目详情的值-废弃
    ClassifyCode   string `json:"classify_code,omitempty" xml:"classify_code,omitempty"`
    // 项目大类名称-取项目详情的值-废弃
    ClassifyName   string `json:"classify_name,omitempty" xml:"classify_name,omitempty"`
    // 国家
    Country   *IdNameDTO `json:"country,omitempty" xml:"country,omitempty"`
    // 项目介绍
    Introduce   string `json:"introduce,omitempty" xml:"introduce,omitempty"`
    // 是否有座： 0=无座 1=有座
    IsHasSeat   int64 `json:"is_has_seat,omitempty" xml:"is_has_seat,omitempty"`
    // 场次
    PerformInfoList   []PerformInfoDTO `json:"perform_info_list,omitempty" xml:"perform_info_list>perform_info_dto,omitempty"`
    // 项目海报图片地址-没用。取项目详情接口的内容
    PosterUrl   string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
    // 项目ID
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 项目名称
    ProjectName   string `json:"project_name,omitempty" xml:"project_name,omitempty"`
    // 项目状态，0：创建中 10：已创建 20：待销售 30：销售中 40：项目取消 50：项目结束 60 定时开售
    ProjectStatus   int64 `json:"project_status,omitempty" xml:"project_status,omitempty"`
    // 项目类型 0:普通项目 1:预售项目 2:先付先抢-先付项目 3:先付先抢-先抢项目 4:搭售项目 5:超级票-暂时没用
    ProjectType   int64 `json:"project_type,omitempty" xml:"project_type,omitempty"`
    // 省
    Province   *IdNameDTO `json:"province,omitempty" xml:"province,omitempty"`
    // 备注-没用
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 项目子类编码-取项目详情的值-废弃
    SubClassifyCode   string `json:"sub_classify_code,omitempty" xml:"sub_classify_code,omitempty"`
    // 项目子类名称-取项目详情的值-废弃
    SubClassifyName   string `json:"sub_classify_name,omitempty" xml:"sub_classify_name,omitempty"`
    // 项目三级分类编码-取项目详情的值-废弃
    ThirdClassifyCode   string `json:"third_classify_code,omitempty" xml:"third_classify_code,omitempty"`
    // 项目三级分类名称-取项目详情的值-废弃
    ThirdClassifyName   string `json:"third_classify_name,omitempty" xml:"third_classify_name,omitempty"`
    // 票务代理费（单位：百分比）
    TicketAgencyFee   string `json:"ticket_agency_fee,omitempty" xml:"ticket_agency_fee,omitempty"`
    // 主办方ID
    TraderIdList   []int64 `json:"trader_id_list,omitempty" xml:"trader_id_list>int64,omitempty"`
    // 主办方名称
    TraderNameList   []string `json:"trader_name_list,omitempty" xml:"trader_name_list>string,omitempty"`
    // 场馆
    Venue   *VenueDTO `json:"venue,omitempty" xml:"venue,omitempty"`
    // 是否总票代
    IsGeneralAgent   int64 `json:"is_general_agent,omitempty" xml:"is_general_agent,omitempty"`
    // 是否测试项目 0-正式项目 1-测试项目
    IsTest   int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
