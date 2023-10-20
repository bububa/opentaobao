package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse 向活动人群实例中增加买家（隐私号版） API返回值
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
type TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse struct {
	model.CommonResponse
	TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponseModel).Reset()
}

// TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponseModel is 向活动人群实例中增加买家（隐私号版） 成功返回结果
type TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_exchange_crowdinstance_add_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	SubSuccess bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubSuccess = false
}

var poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse)
	},
}

// GetTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse 从 sync.Pool 获取 TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse
func GetTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse() *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse {
	return poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse.Get().(*TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse)
}

// ReleaseTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse 将 TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse(v *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse) {
	v.Reset()
	poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIResponse.Put(v)
}
