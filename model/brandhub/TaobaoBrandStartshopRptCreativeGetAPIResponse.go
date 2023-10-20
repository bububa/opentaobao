package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptCreativeGetAPIResponse 明星店铺创意报表数据查询 API返回值
// taobao.brand.startshop.rpt.creative.get
//
// 获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptCreativeGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStartshopRptCreativeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptCreativeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandStartshopRptCreativeGetAPIResponseModel).Reset()
}

// TaobaoBrandStartshopRptCreativeGetAPIResponseModel is 明星店铺创意报表数据查询 成功返回结果
type TaobaoBrandStartshopRptCreativeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_startshop_rpt_creative_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStartshopRptCreativeGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_creative_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptCreativeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignRptList = m.CampaignRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandStartshopRptCreativeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStartshopRptCreativeGetAPIResponse)
	},
}

// GetTaobaoBrandStartshopRptCreativeGetAPIResponse 从 sync.Pool 获取 TaobaoBrandStartshopRptCreativeGetAPIResponse
func GetTaobaoBrandStartshopRptCreativeGetAPIResponse() *TaobaoBrandStartshopRptCreativeGetAPIResponse {
	return poolTaobaoBrandStartshopRptCreativeGetAPIResponse.Get().(*TaobaoBrandStartshopRptCreativeGetAPIResponse)
}

// ReleaseTaobaoBrandStartshopRptCreativeGetAPIResponse 将 TaobaoBrandStartshopRptCreativeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandStartshopRptCreativeGetAPIResponse(v *TaobaoBrandStartshopRptCreativeGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandStartshopRptCreativeGetAPIResponse.Put(v)
}
