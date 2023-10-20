package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityFatigueAPIResponse 营销抽奖-弹窗疲劳度控制 API返回值
// alitrip.merchant.galaxy.activity.fatigue
//
// 星河产品-营销抽奖首页弹窗疲劳度控制服务
type AlitripMerchantGalaxyActivityFatigueAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityFatigueAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityFatigueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyActivityFatigueAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyActivityFatigueAPIResponseModel is 营销抽奖-弹窗疲劳度控制 成功返回结果
type AlitripMerchantGalaxyActivityFatigueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_fatigue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityFatigueResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityFatigueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyActivityFatigueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityFatigueAPIResponse)
	},
}

// GetAlitripMerchantGalaxyActivityFatigueAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyActivityFatigueAPIResponse
func GetAlitripMerchantGalaxyActivityFatigueAPIResponse() *AlitripMerchantGalaxyActivityFatigueAPIResponse {
	return poolAlitripMerchantGalaxyActivityFatigueAPIResponse.Get().(*AlitripMerchantGalaxyActivityFatigueAPIResponse)
}

// ReleaseAlitripMerchantGalaxyActivityFatigueAPIResponse 将 AlitripMerchantGalaxyActivityFatigueAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityFatigueAPIResponse(v *AlitripMerchantGalaxyActivityFatigueAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityFatigueAPIResponse.Put(v)
}
