package brandhub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺定向维度报表 API返回值 
taobao.brand.starshop.rpt.target.get

获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStarshopRptTargetGetAPIResponse struct {
    model.CommonResponse
    TaobaoBrandStarshopRptTargetGetResponse
}

// 明星店铺定向维度报表 成功返回结果
type TaobaoBrandStarshopRptTargetGetResponse struct {
    XMLName xml.Name `xml:"brand_starshop_rpt_target_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrorParam   string `json:"error_param,omitempty" xml:"error_param,omitempty"`
    // 返回结果
    CampaignRptList   []TaobaoBrandStarshopRptTargetGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_starshop_rpt_target_get_result,omitempty"`
}
