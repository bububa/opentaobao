package wangwang

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWangwangAbstractAddwordAPIResponse 增加关键词 API返回值
// taobao.wangwang.abstract.addword
//
// 增加关键词，只支持json返回
type TaobaoWangwangAbstractAddwordAPIResponse struct {
	model.CommonResponse
	TaobaoWangwangAbstractAddwordAPIResponseModel
}

// TaobaoWangwangAbstractAddwordAPIResponseModel is 增加关键词 成功返回结果
type TaobaoWangwangAbstractAddwordAPIResponseModel struct {
	XMLName xml.Name `xml:"wangwang_abstract_addword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 例如单词长度太长等，当ret_code=-1时才有这项
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 0或-1，表示错误或正确，错误时有错误信息
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
