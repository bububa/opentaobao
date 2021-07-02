package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstShiporderCancelAPIResponse 零售通发货单取消 API返回值
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
type AlibabaLstShiporderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaLstShiporderCancelAPIResponseModel
}

// AlibabaLstShiporderCancelAPIResponseModel is 零售通发货单取消 成功返回结果
type AlibabaLstShiporderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_shiporder_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
