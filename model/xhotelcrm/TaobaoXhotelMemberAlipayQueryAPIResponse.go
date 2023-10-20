package xhotelcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelmemberalipayqueryAPIResponse 希尔顿会员查询 API返回值
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
type TaobaoxhotelmemberalipayqueryAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelmemberalipayqueryAPIResponseModel
}

// TaobaoxhotelmemberalipayqueryAPIResponseModel is 希尔顿会员查询 成功返回结果
type TaobaoxhotelmemberalipayqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_member_alipay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HmemberResult `json:"result,omitempty" xml:"result,omitempty"`
}
