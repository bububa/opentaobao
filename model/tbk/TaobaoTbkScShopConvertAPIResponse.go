package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScShopConvertAPIResponse 淘宝客-服务商-店铺链接转换 API返回值
// taobao.tbk.sc.shop.convert
//
// 服务商使用。支持入参推广者对应的“推广位”和卖家id，获取对应的店铺推广链接。
type TaobaoTbkScShopConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScShopConvertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScShopConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScShopConvertAPIResponseModel).Reset()
}

// TaobaoTbkScShopConvertAPIResponseModel is 淘宝客-服务商-店铺链接转换 成功返回结果
type TaobaoTbkScShopConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_shop_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加入淘宝客的店铺
	Results []NTbkShop `json:"results,omitempty" xml:"results>n_tbk_shop,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScShopConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoTbkScShopConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScShopConvertAPIResponse)
	},
}

// GetTaobaoTbkScShopConvertAPIResponse 从 sync.Pool 获取 TaobaoTbkScShopConvertAPIResponse
func GetTaobaoTbkScShopConvertAPIResponse() *TaobaoTbkScShopConvertAPIResponse {
	return poolTaobaoTbkScShopConvertAPIResponse.Get().(*TaobaoTbkScShopConvertAPIResponse)
}

// ReleaseTaobaoTbkScShopConvertAPIResponse 将 TaobaoTbkScShopConvertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScShopConvertAPIResponse(v *TaobaoTbkScShopConvertAPIResponse) {
	v.Reset()
	poolTaobaoTbkScShopConvertAPIResponse.Put(v)
}
