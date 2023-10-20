package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse 储值账户充值前校验 API返回值
// alibaba.alsc.crm.recharge.chargeprecheck.get
//
// 储值账户充值前校验接口
type AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel is 储值账户充值前校验 成功返回结果
type AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_chargeprecheck_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeChargeprecheckGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse
func GetAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse() *AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse {
	return poolAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse.Get().(*AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse 将 AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse(v *AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeChargeprecheckGetAPIResponse.Put(v)
}
