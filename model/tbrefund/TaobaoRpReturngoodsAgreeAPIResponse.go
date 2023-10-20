package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsAgreeAPIResponse 卖家同意退货 API返回值
// taobao.rp.returngoods.agree
//
// 卖家同意退货，支持淘宝和天猫的订单。
type TaobaoRpReturngoodsAgreeAPIResponse struct {
	model.CommonResponse
	TaobaoRpReturngoodsAgreeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsAgreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRpReturngoodsAgreeAPIResponseModel).Reset()
}

// TaobaoRpReturngoodsAgreeAPIResponseModel is 卖家同意退货 成功返回结果
type TaobaoRpReturngoodsAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_returngoods_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsAgreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoRpReturngoodsAgreeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRpReturngoodsAgreeAPIResponse)
	},
}

// GetTaobaoRpReturngoodsAgreeAPIResponse 从 sync.Pool 获取 TaobaoRpReturngoodsAgreeAPIResponse
func GetTaobaoRpReturngoodsAgreeAPIResponse() *TaobaoRpReturngoodsAgreeAPIResponse {
	return poolTaobaoRpReturngoodsAgreeAPIResponse.Get().(*TaobaoRpReturngoodsAgreeAPIResponse)
}

// ReleaseTaobaoRpReturngoodsAgreeAPIResponse 将 TaobaoRpReturngoodsAgreeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRpReturngoodsAgreeAPIResponse(v *TaobaoRpReturngoodsAgreeAPIResponse) {
	v.Reset()
	poolTaobaoRpReturngoodsAgreeAPIResponse.Put(v)
}
