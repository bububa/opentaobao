package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniussendgoodscancelAPIResponse 取消发货 API返回值
// taobao.rdc.aligenius.sendgoods.cancel
//
// 提供商家在仅退款中发送取消发货状态
type TaobaordcaligeniussendgoodscancelAPIResponse struct {
	model.CommonResponse
	TaobaordcaligeniussendgoodscancelAPIResponseModel
}

// TaobaordcaligeniussendgoodscancelAPIResponseModel is 取消发货 成功返回结果
type TaobaordcaligeniussendgoodscancelAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_sendgoods_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaordcaligeniussendgoodscancelResult `json:"result,omitempty" xml:"result,omitempty"`
}
