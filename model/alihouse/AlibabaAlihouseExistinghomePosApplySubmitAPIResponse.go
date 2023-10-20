package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeposapplysubmitAPIResponse pos申请接口 API返回值
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
type AlibabaalihouseexistinghomeposapplysubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeposapplysubmitAPIResponseModel
}

// AlibabaalihouseexistinghomeposapplysubmitAPIResponseModel is pos申请接口 成功返回结果
type AlibabaalihouseexistinghomeposapplysubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_pos_apply_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象
	Result *AlibabaalihouseexistinghomeposapplysubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
