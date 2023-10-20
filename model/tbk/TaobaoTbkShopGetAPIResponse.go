package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkShopGetAPIResponse 淘宝客-推广者-店铺搜索 API返回值
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
type TaobaoTbkShopGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkShopGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkShopGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkShopGetAPIResponseModel).Reset()
}

// TaobaoTbkShopGetAPIResponseModel is 淘宝客-推广者-店铺搜索 成功返回结果
type TaobaoTbkShopGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝客店铺
	Results []NTbkShop `json:"results,omitempty" xml:"results>n_tbk_shop,omitempty"`
	// 搜索到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkShopGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalResults = 0
}

var poolTaobaoTbkShopGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkShopGetAPIResponse)
	},
}

// GetTaobaoTbkShopGetAPIResponse 从 sync.Pool 获取 TaobaoTbkShopGetAPIResponse
func GetTaobaoTbkShopGetAPIResponse() *TaobaoTbkShopGetAPIResponse {
	return poolTaobaoTbkShopGetAPIResponse.Get().(*TaobaoTbkShopGetAPIResponse)
}

// ReleaseTaobaoTbkShopGetAPIResponse 将 TaobaoTbkShopGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkShopGetAPIResponse(v *TaobaoTbkShopGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkShopGetAPIResponse.Put(v)
}
