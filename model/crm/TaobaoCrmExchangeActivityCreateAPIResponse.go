package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeActivityCreateAPIResponse 创建积分兑换活动 API返回值
// taobao.crm.exchange.activity.create
//
// 创建针对积分兑换类型的活动
type TaobaoCrmExchangeActivityCreateAPIResponse struct {
	model.CommonResponse
	TaobaoCrmExchangeActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmExchangeActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmExchangeActivityCreateAPIResponseModel).Reset()
}

// TaobaoCrmExchangeActivityCreateAPIResponseModel is 创建积分兑换活动 成功返回结果
type TaobaoCrmExchangeActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_exchange_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 人群实例ID
	CrowdinstanceId int64 `json:"crowdinstance_id,omitempty" xml:"crowdinstance_id,omitempty"`
	// 接口调用成功
	SubSuccess bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmExchangeActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ActivityId = 0
	m.CrowdinstanceId = 0
	m.SubSuccess = false
}

var poolTaobaoCrmExchangeActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmExchangeActivityCreateAPIResponse)
	},
}

// GetTaobaoCrmExchangeActivityCreateAPIResponse 从 sync.Pool 获取 TaobaoCrmExchangeActivityCreateAPIResponse
func GetTaobaoCrmExchangeActivityCreateAPIResponse() *TaobaoCrmExchangeActivityCreateAPIResponse {
	return poolTaobaoCrmExchangeActivityCreateAPIResponse.Get().(*TaobaoCrmExchangeActivityCreateAPIResponse)
}

// ReleaseTaobaoCrmExchangeActivityCreateAPIResponse 将 TaobaoCrmExchangeActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmExchangeActivityCreateAPIResponse(v *TaobaoCrmExchangeActivityCreateAPIResponse) {
	v.Reset()
	poolTaobaoCrmExchangeActivityCreateAPIResponse.Put(v)
}
