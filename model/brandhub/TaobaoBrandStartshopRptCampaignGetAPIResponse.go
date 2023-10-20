package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptCampaignGetAPIResponse 明星店铺推广计划报表数据查询 API返回值
// taobao.brand.startshop.rpt.campaign.get
//
// 获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptCampaignGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStartshopRptCampaignGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptCampaignGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandStartshopRptCampaignGetAPIResponseModel).Reset()
}

// TaobaoBrandStartshopRptCampaignGetAPIResponseModel is 明星店铺推广计划报表数据查询 成功返回结果
type TaobaoBrandStartshopRptCampaignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_startshop_rpt_campaign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStartshopRptCampaignGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_campaign_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptCampaignGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignRptList = m.CampaignRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandStartshopRptCampaignGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStartshopRptCampaignGetAPIResponse)
	},
}

// GetTaobaoBrandStartshopRptCampaignGetAPIResponse 从 sync.Pool 获取 TaobaoBrandStartshopRptCampaignGetAPIResponse
func GetTaobaoBrandStartshopRptCampaignGetAPIResponse() *TaobaoBrandStartshopRptCampaignGetAPIResponse {
	return poolTaobaoBrandStartshopRptCampaignGetAPIResponse.Get().(*TaobaoBrandStartshopRptCampaignGetAPIResponse)
}

// ReleaseTaobaoBrandStartshopRptCampaignGetAPIResponse 将 TaobaoBrandStartshopRptCampaignGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandStartshopRptCampaignGetAPIResponse(v *TaobaoBrandStartshopRptCampaignGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandStartshopRptCampaignGetAPIResponse.Put(v)
}
