package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeposopensubmitAPIResponse pos进件接口 API返回值
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
type AlibabaalihouseexistinghomeposopensubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeposopensubmitAPIResponseModel
}

// AlibabaalihouseexistinghomeposopensubmitAPIResponseModel is pos进件接口 成功返回结果
type AlibabaalihouseexistinghomeposopensubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_pos_open_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaalihouseexistinghomeposopensubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
