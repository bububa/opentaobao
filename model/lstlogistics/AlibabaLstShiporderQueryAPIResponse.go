package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstshiporderqueryAPIResponse 零售通发货单查询 API返回值
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
type AlibabalstshiporderqueryAPIResponse struct {
	model.CommonResponse
	AlibabalstshiporderqueryAPIResponseModel
}

// AlibabalstshiporderqueryAPIResponseModel is 零售通发货单查询 成功返回结果
type AlibabalstshiporderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_shiporder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
