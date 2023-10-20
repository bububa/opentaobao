package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsRefuseAPIResponse 卖家拒绝退货 API返回值
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
type TaobaoRpReturngoodsRefuseAPIResponse struct {
	model.CommonResponse
	TaobaoRpReturngoodsRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRpReturngoodsRefuseAPIResponseModel).Reset()
}

// TaobaoRpReturngoodsRefuseAPIResponseModel is 卖家拒绝退货 成功返回结果
type TaobaoRpReturngoodsRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"rp_returngoods_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// asdf
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRpReturngoodsRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoRpReturngoodsRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRpReturngoodsRefuseAPIResponse)
	},
}

// GetTaobaoRpReturngoodsRefuseAPIResponse 从 sync.Pool 获取 TaobaoRpReturngoodsRefuseAPIResponse
func GetTaobaoRpReturngoodsRefuseAPIResponse() *TaobaoRpReturngoodsRefuseAPIResponse {
	return poolTaobaoRpReturngoodsRefuseAPIResponse.Get().(*TaobaoRpReturngoodsRefuseAPIResponse)
}

// ReleaseTaobaoRpReturngoodsRefuseAPIResponse 将 TaobaoRpReturngoodsRefuseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRpReturngoodsRefuseAPIResponse(v *TaobaoRpReturngoodsRefuseAPIResponse) {
	v.Reset()
	poolTaobaoRpReturngoodsRefuseAPIResponse.Put(v)
}
