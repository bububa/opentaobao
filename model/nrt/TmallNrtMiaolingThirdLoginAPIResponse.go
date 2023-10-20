package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtmiaolingthirdloginAPIResponse 喵零第三方免登 API返回值
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
type TmallnrtmiaolingthirdloginAPIResponse struct {
	model.CommonResponse
	TmallnrtmiaolingthirdloginAPIResponseModel
}

// TmallnrtmiaolingthirdloginAPIResponseModel is 喵零第三方免登 成功返回结果
type TmallnrtmiaolingthirdloginAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_miaoling_third_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
