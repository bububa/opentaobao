package xhotelcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberAlipayQueryAPIResponse 希尔顿会员查询 API返回值
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
type TaobaoXhotelMemberAlipayQueryAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMemberAlipayQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelMemberAlipayQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelMemberAlipayQueryAPIResponseModel).Reset()
}

// TaobaoXhotelMemberAlipayQueryAPIResponseModel is 希尔顿会员查询 成功返回结果
type TaobaoXhotelMemberAlipayQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_member_alipay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HMemberResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelMemberAlipayQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelMemberAlipayQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelMemberAlipayQueryAPIResponse)
	},
}

// GetTaobaoXhotelMemberAlipayQueryAPIResponse 从 sync.Pool 获取 TaobaoXhotelMemberAlipayQueryAPIResponse
func GetTaobaoXhotelMemberAlipayQueryAPIResponse() *TaobaoXhotelMemberAlipayQueryAPIResponse {
	return poolTaobaoXhotelMemberAlipayQueryAPIResponse.Get().(*TaobaoXhotelMemberAlipayQueryAPIResponse)
}

// ReleaseTaobaoXhotelMemberAlipayQueryAPIResponse 将 TaobaoXhotelMemberAlipayQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelMemberAlipayQueryAPIResponse(v *TaobaoXhotelMemberAlipayQueryAPIResponse) {
	v.Reset()
	poolTaobaoXhotelMemberAlipayQueryAPIResponse.Put(v)
}
