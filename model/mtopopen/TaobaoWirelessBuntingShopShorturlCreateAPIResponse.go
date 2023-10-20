package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessBuntingShopShorturlCreateAPIResponse 通过店铺id取得短链 API返回值
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
type TaobaoWirelessBuntingShopShorturlCreateAPIResponse struct {
	model.CommonResponse
	TaobaoWirelessBuntingShopShorturlCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWirelessBuntingShopShorturlCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWirelessBuntingShopShorturlCreateAPIResponseModel).Reset()
}

// TaobaoWirelessBuntingShopShorturlCreateAPIResponseModel is 通过店铺id取得短链 成功返回结果
type TaobaoWirelessBuntingShopShorturlCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"wireless_bunting_shop_shorturl_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短链
	Shorturl string `json:"shorturl,omitempty" xml:"shorturl,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWirelessBuntingShopShorturlCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shorturl = ""
}

var poolTaobaoWirelessBuntingShopShorturlCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWirelessBuntingShopShorturlCreateAPIResponse)
	},
}

// GetTaobaoWirelessBuntingShopShorturlCreateAPIResponse 从 sync.Pool 获取 TaobaoWirelessBuntingShopShorturlCreateAPIResponse
func GetTaobaoWirelessBuntingShopShorturlCreateAPIResponse() *TaobaoWirelessBuntingShopShorturlCreateAPIResponse {
	return poolTaobaoWirelessBuntingShopShorturlCreateAPIResponse.Get().(*TaobaoWirelessBuntingShopShorturlCreateAPIResponse)
}

// ReleaseTaobaoWirelessBuntingShopShorturlCreateAPIResponse 将 TaobaoWirelessBuntingShopShorturlCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWirelessBuntingShopShorturlCreateAPIResponse(v *TaobaoWirelessBuntingShopShorturlCreateAPIResponse) {
	v.Reset()
	poolTaobaoWirelessBuntingShopShorturlCreateAPIResponse.Put(v)
}
