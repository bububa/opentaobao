package brandhub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺推广计划报表数据查询 API返回值 
taobao.brand.startshop.rpt.campaign.get

获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptCampaignGetAPIResponse struct {
    model.CommonResponse
    TaobaoBrandStartshopRptCampaignGetAPIResponseModel
}

// 明星店铺推广计划报表数据查询 成功返回结果
type TaobaoBrandStartshopRptCampaignGetAPIResponseModel struct {
    XMLName xml.Name `xml:"brand_startshop_rpt_campaign_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrorParam   string `json:"error_param,omitempty" xml:"error_param,omitempty"`
    // 返回结果
    CampaignRptList   []TaobaoBrandStartshopRptCampaignGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_campaign_get_result,omitempty"`
}
