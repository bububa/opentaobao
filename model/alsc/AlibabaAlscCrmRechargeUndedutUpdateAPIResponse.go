package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUndedutUpdateAPIResponse 储值消费退款(逆向) API返回值
// alibaba.alsc.crm.recharge.undedut.update
//
// 新增储值消费退款接口
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeUndedutUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel is 储值消费退款(逆向) 成功返回结果
type AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_undedut_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeUndedutUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeUndedutUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeUndedutUpdateAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeUndedutUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeUndedutUpdateAPIResponse
func GetAlibabaAlscCrmRechargeUndedutUpdateAPIResponse() *AlibabaAlscCrmRechargeUndedutUpdateAPIResponse {
	return poolAlibabaAlscCrmRechargeUndedutUpdateAPIResponse.Get().(*AlibabaAlscCrmRechargeUndedutUpdateAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeUndedutUpdateAPIResponse 将 AlibabaAlscCrmRechargeUndedutUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeUndedutUpdateAPIResponse(v *AlibabaAlscCrmRechargeUndedutUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeUndedutUpdateAPIResponse.Put(v)
}
