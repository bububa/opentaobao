package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeChargeUpdateAPIResponse 储值充值 API返回值
// alibaba.alsc.crm.recharge.charge.update
//
// 顾客储值账户充值
type AlibabaAlscCrmRechargeChargeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeChargeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeChargeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeChargeUpdateAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeChargeUpdateAPIResponseModel is 储值充值 成功返回结果
type AlibabaAlscCrmRechargeChargeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_charge_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeChargeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeChargeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeChargeUpdateAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeChargeUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeChargeUpdateAPIResponse
func GetAlibabaAlscCrmRechargeChargeUpdateAPIResponse() *AlibabaAlscCrmRechargeChargeUpdateAPIResponse {
	return poolAlibabaAlscCrmRechargeChargeUpdateAPIResponse.Get().(*AlibabaAlscCrmRechargeChargeUpdateAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeChargeUpdateAPIResponse 将 AlibabaAlscCrmRechargeChargeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeChargeUpdateAPIResponse(v *AlibabaAlscCrmRechargeChargeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeChargeUpdateAPIResponse.Put(v)
}
