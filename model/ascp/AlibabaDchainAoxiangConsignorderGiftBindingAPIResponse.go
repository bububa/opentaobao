package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignordergiftbindingAPIResponse 赠品绑赠计算占用 API返回值
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
type AlibabadchainaoxiangconsignordergiftbindingAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangconsignordergiftbindingAPIResponseModel
}

// AlibabadchainaoxiangconsignordergiftbindingAPIResponseModel is 赠品绑赠计算占用 成功返回结果
type AlibabadchainaoxiangconsignordergiftbindingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_gift_binding_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BindingConsignorderGiftResponse *BindingConsignOrderGiftRequest `json:"binding_consignorder_gift_response,omitempty" xml:"binding_consignorder_gift_response,omitempty"`
}
