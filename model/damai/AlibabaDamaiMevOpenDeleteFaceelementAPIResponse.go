package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeletefaceelementAPIResponse 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement API返回值
// alibaba.damai.mev.open.delete.faceelement
//
// deleteFaceElement
type AlibabadamaimevopendeletefaceelementAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopendeletefaceelementAPIResponseModel
}

// AlibabadamaimevopendeletefaceelementAPIResponseModel is 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement 成功返回结果
type AlibabadamaimevopendeletefaceelementAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_delete_faceelement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopendeletefaceelementResult `json:"result,omitempty" xml:"result,omitempty"`
}
