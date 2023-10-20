package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignsGetAPIResponse 取得一组推广计划 API返回值
// taobao.simba.campaigns.get
//
// 取得一个客户的推广计划；
type TaobaoSimbaCampaignsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignsGetAPIResponseModel
}

// TaobaoSimbaCampaignsGetAPIResponseModel is 取得一组推广计划 成功返回结果
type TaobaoSimbaCampaignsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaigns_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划列表
	Campaigns []Campaign `json:"campaigns,omitempty" xml:"campaigns>campaign,omitempty"`
}
