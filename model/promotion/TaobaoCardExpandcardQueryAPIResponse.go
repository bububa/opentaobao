package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocardexpandcardqueryAPIResponse 购物金卡查询 API返回值
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
type TaobaocardexpandcardqueryAPIResponse struct {
	model.CommonResponse
	TaobaocardexpandcardqueryAPIResponseModel
}

// TaobaocardexpandcardqueryAPIResponseModel is 购物金卡查询 成功返回结果
type TaobaocardexpandcardqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"card_expandcard_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaocardexpandcardqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
