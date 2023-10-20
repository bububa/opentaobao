package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillPrivacySubscriptionGetAPIResponse 隐私面单商家订购查询 API返回值
// cainiao.waybill.privacy.subscription.get
//
// ISV查询商家是否订购隐私面单
type CainiaoWaybillPrivacySubscriptionGetAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillPrivacySubscriptionGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillPrivacySubscriptionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillPrivacySubscriptionGetAPIResponseModel).Reset()
}

// CainiaoWaybillPrivacySubscriptionGetAPIResponseModel is 隐私面单商家订购查询 成功返回结果
type CainiaoWaybillPrivacySubscriptionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_privacy_subscription_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *CainiaoWaybillPrivacySubscriptionGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillPrivacySubscriptionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoWaybillPrivacySubscriptionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillPrivacySubscriptionGetAPIResponse)
	},
}

// GetCainiaoWaybillPrivacySubscriptionGetAPIResponse 从 sync.Pool 获取 CainiaoWaybillPrivacySubscriptionGetAPIResponse
func GetCainiaoWaybillPrivacySubscriptionGetAPIResponse() *CainiaoWaybillPrivacySubscriptionGetAPIResponse {
	return poolCainiaoWaybillPrivacySubscriptionGetAPIResponse.Get().(*CainiaoWaybillPrivacySubscriptionGetAPIResponse)
}

// ReleaseCainiaoWaybillPrivacySubscriptionGetAPIResponse 将 CainiaoWaybillPrivacySubscriptionGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillPrivacySubscriptionGetAPIResponse(v *CainiaoWaybillPrivacySubscriptionGetAPIResponse) {
	v.Reset()
	poolCainiaoWaybillPrivacySubscriptionGetAPIResponse.Put(v)
}
