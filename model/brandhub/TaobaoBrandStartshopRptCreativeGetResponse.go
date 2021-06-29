package brandhub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺创意报表数据查询 APIResponse
taobao.brand.startshop.rpt.creative.get

获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptCreativeGetAPIResponse struct {
    model.CommonResponse
    TaobaoBrandStartshopRptCreativeGetResponse
}

type TaobaoBrandStartshopRptCreativeGetResponse struct {
    XMLName xml.Name `xml:"brand_startshop_rpt_creative_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    ErrorParam   string `json:"error_param,omitempty" xml:"error_param,omitempty"`

    
    // 返回结果
    
    CampaignRptList   []TaobaoBrandStartshopRptCreativeGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_creative_get_result,omitempty"`
    
    
}
