package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmsfverifyAPIResponse 喵师傅核销接口 API返回值
// tmall.msf.verify
//
// msf服务核销的top接口
type TmallmsfverifyAPIResponse struct {
	model.CommonResponse
	TmallmsfverifyAPIResponseModel
}

// TmallmsfverifyAPIResponseModel is 喵师傅核销接口 成功返回结果
type TmallmsfverifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
