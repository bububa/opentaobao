package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenrechargeoperateAPIResponse 储值操作接口 API返回值
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
type AlibabaalsccrmopenrechargeoperateAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmopenrechargeoperateAPIResponseModel
}

// AlibabaalsccrmopenrechargeoperateAPIResponseModel is 储值操作接口 成功返回结果
type AlibabaalsccrmopenrechargeoperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_recharge_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
