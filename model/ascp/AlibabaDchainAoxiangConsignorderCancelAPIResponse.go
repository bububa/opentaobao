package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignordercancelAPIResponse 自动流转单据取消仓内发货 API返回值
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
type AlibabadchainaoxiangconsignordercancelAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangconsignordercancelAPIResponseModel
}

// AlibabadchainaoxiangconsignordercancelAPIResponseModel is 自动流转单据取消仓内发货 成功返回结果
type AlibabadchainaoxiangconsignordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CancelConsignorderResponse *CancelConsignOrderResponse `json:"cancel_consignorder_response,omitempty" xml:"cancel_consignorder_response,omitempty"`
}
