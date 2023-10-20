package sungari

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosungaridisposequeryAPIResponse 商品商家处置结果查询 API返回值
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
type TaobaosungaridisposequeryAPIResponse struct {
	model.CommonResponse
	TaobaosungaridisposequeryAPIResponseModel
}

// TaobaosungaridisposequeryAPIResponseModel is 商品商家处置结果查询 成功返回结果
type TaobaosungaridisposequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"sungari_dispose_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosungaridisposequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
