package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuBuyerTagGetAPIResponse 判断买家是否有某些标 API返回值
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
type TaobaoQianniuBuyerTagGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuBuyerTagGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuBuyerTagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuBuyerTagGetAPIResponseModel).Reset()
}

// TaobaoQianniuBuyerTagGetAPIResponseModel is 判断买家是否有某些标 成功返回结果
type TaobaoQianniuBuyerTagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_buyer_tag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户tag信息
	UserTagInfo *UserTagQueryResult `json:"user_tag_info,omitempty" xml:"user_tag_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuBuyerTagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UserTagInfo = nil
}

var poolTaobaoQianniuBuyerTagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuBuyerTagGetAPIResponse)
	},
}

// GetTaobaoQianniuBuyerTagGetAPIResponse 从 sync.Pool 获取 TaobaoQianniuBuyerTagGetAPIResponse
func GetTaobaoQianniuBuyerTagGetAPIResponse() *TaobaoQianniuBuyerTagGetAPIResponse {
	return poolTaobaoQianniuBuyerTagGetAPIResponse.Get().(*TaobaoQianniuBuyerTagGetAPIResponse)
}

// ReleaseTaobaoQianniuBuyerTagGetAPIResponse 将 TaobaoQianniuBuyerTagGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuBuyerTagGetAPIResponse(v *TaobaoQianniuBuyerTagGetAPIResponse) {
	v.Reset()
	poolTaobaoQianniuBuyerTagGetAPIResponse.Put(v)
}
