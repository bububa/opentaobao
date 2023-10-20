package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfReceiveAPIResponse 签收接口 API返回值
// tmall.msf.receive
//
// 签收接口
type TmallMsfReceiveAPIResponse struct {
	model.CommonResponse
	TmallMsfReceiveAPIResponseModel
}

// TmallMsfReceiveAPIResponseModel is 签收接口 成功返回结果
type TmallMsfReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
