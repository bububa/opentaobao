package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReceiverinfoQueryAPIResponse OAID 收件人信息解密接口 API返回值
// taobao.qimen.receiverinfo.query
//
// WMS 调用该接口，通过 OAID 查询解密后的收件人信息
type TaobaoQimenReceiverinfoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReceiverinfoQueryAPIResponseModel
}

// TaobaoQimenReceiverinfoQueryAPIResponseModel is OAID 收件人信息解密接口 成功返回结果
type TaobaoQimenReceiverinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_receiverinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReceiverinfoQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
