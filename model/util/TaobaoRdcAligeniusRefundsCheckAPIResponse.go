package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusrefundscheckAPIResponse 退款信息审核 API返回值
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
type TaobaordcaligeniusrefundscheckAPIResponse struct {
	model.CommonResponse
	TaobaordcaligeniusrefundscheckAPIResponseModel
}

// TaobaordcaligeniusrefundscheckAPIResponseModel is 退款信息审核 成功返回结果
type TaobaordcaligeniusrefundscheckAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_refunds_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordcaligeniusrefundscheckResult `json:"result,omitempty" xml:"result,omitempty"`
}
