package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopOnceTokenGetAPIResponse
网关一次性token获取 API返回值
taobao.top.once.token.get

网关一次性token获取 */
type TaobaoTopOnceTokenGetAPIResponse struct {
	model.CommonResponse
	TaobaoTopOnceTokenGetAPIResponseModel
}

// TaobaoTopOnceTokenGetAPIResponseModel is 网关一次性token获取 成功返回结果
type TaobaoTopOnceTokenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_once_token_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 响应编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 失败详情
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
