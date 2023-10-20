package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptWordpackageGetAPIResponse 明星店铺品牌流量包报表数据查询 API返回值
// taobao.brand.startshop.rpt.wordpackage.get
//
// 获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptWordpackageGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStartshopRptWordpackageGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptWordpackageGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandStartshopRptWordpackageGetAPIResponseModel).Reset()
}

// TaobaoBrandStartshopRptWordpackageGetAPIResponseModel is 明星店铺品牌流量包报表数据查询 成功返回结果
type TaobaoBrandStartshopRptWordpackageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_startshop_rpt_wordpackage_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStartshopRptWordpackageGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_wordpackage_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptWordpackageGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignRptList = m.CampaignRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandStartshopRptWordpackageGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStartshopRptWordpackageGetAPIResponse)
	},
}

// GetTaobaoBrandStartshopRptWordpackageGetAPIResponse 从 sync.Pool 获取 TaobaoBrandStartshopRptWordpackageGetAPIResponse
func GetTaobaoBrandStartshopRptWordpackageGetAPIResponse() *TaobaoBrandStartshopRptWordpackageGetAPIResponse {
	return poolTaobaoBrandStartshopRptWordpackageGetAPIResponse.Get().(*TaobaoBrandStartshopRptWordpackageGetAPIResponse)
}

// ReleaseTaobaoBrandStartshopRptWordpackageGetAPIResponse 将 TaobaoBrandStartshopRptWordpackageGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandStartshopRptWordpackageGetAPIResponse(v *TaobaoBrandStartshopRptWordpackageGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandStartshopRptWordpackageGetAPIResponse.Put(v)
}
