package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallSscWorkcardAcceptAPIResponse 服务商接单完成 API返回值
// tmall.ssc.workcard.accept
//
// 工单完结
type TmallSscWorkcardAcceptAPIResponse struct {
	model.CommonResponse
	TmallSscWorkcardAcceptAPIResponseModel
}

// TmallSscWorkcardAcceptAPIResponseModel is 服务商接单完成 成功返回结果
type TmallSscWorkcardAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ssc_workcard_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接单结果
	Result *TmallSscWorkcardAcceptResult `json:"result,omitempty" xml:"result,omitempty"`
}
