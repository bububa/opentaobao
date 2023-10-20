package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppipGetAPIResponse 获取ISV发起请求服务器IP API返回值
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
type TaobaoAppipGetAPIResponse struct {
	model.CommonResponse
	TaobaoAppipGetAPIResponseModel
}

// TaobaoAppipGetAPIResponseModel is 获取ISV发起请求服务器IP 成功返回结果
type TaobaoAppipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"appip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ISV发起请求服务器IP
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
}
