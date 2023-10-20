package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsRefillAPIResponse 卖家回填物流信息 API返回值
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillAPIResponse struct {
	model.CommonResponse
	TaobaoRpReturngoodsRefillAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsRefillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRpReturngoodsRefillAPIResponseModel).Reset()
}

// TaobaoRpReturngoodsRefillAPIResponseModel is 卖家回填物流信息 成功返回结果
type TaobaoRpReturngoodsRefillAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_returngoods_refill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 验货操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsRefillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoRpReturngoodsRefillAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRpReturngoodsRefillAPIResponse)
	},
}

// GetTaobaoRpReturngoodsRefillAPIResponse 从 sync.Pool 获取 TaobaoRpReturngoodsRefillAPIResponse
func GetTaobaoRpReturngoodsRefillAPIResponse() *TaobaoRpReturngoodsRefillAPIResponse {
	return poolTaobaoRpReturngoodsRefillAPIResponse.Get().(*TaobaoRpReturngoodsRefillAPIResponse)
}

// ReleaseTaobaoRpReturngoodsRefillAPIResponse 将 TaobaoRpReturngoodsRefillAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRpReturngoodsRefillAPIResponse(v *TaobaoRpReturngoodsRefillAPIResponse) {
	v.Reset()
	poolTaobaoRpReturngoodsRefillAPIResponse.Put(v)
}
