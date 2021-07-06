package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyaddressOpenAPIResponse 淘宝自助修改地址服务开通 API返回值
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
type TaobaoModifyaddressOpenAPIResponse struct {
	model.CommonResponse
	TaobaoModifyaddressOpenAPIResponseModel
}

// TaobaoModifyaddressOpenAPIResponseModel is 淘宝自助修改地址服务开通 成功返回结果
type TaobaoModifyaddressOpenAPIResponseModel struct {
	XMLName xml.Name `xml:"modifyaddress_open_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
