package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoShopUpdateAPIResponse 更新店铺基本信息 API返回值
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
type TaobaoShopUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoShopUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoShopUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoShopUpdateAPIResponseModel).Reset()
}

// TaobaoShopUpdateAPIResponseModel is 更新店铺基本信息 成功返回结果
type TaobaoShopUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"shop_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺信息
	Shop *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoShopUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Shop = nil
}

var poolTaobaoShopUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoShopUpdateAPIResponse)
	},
}

// GetTaobaoShopUpdateAPIResponse 从 sync.Pool 获取 TaobaoShopUpdateAPIResponse
func GetTaobaoShopUpdateAPIResponse() *TaobaoShopUpdateAPIResponse {
	return poolTaobaoShopUpdateAPIResponse.Get().(*TaobaoShopUpdateAPIResponse)
}

// ReleaseTaobaoShopUpdateAPIResponse 将 TaobaoShopUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoShopUpdateAPIResponse(v *TaobaoShopUpdateAPIResponse) {
	v.Reset()
	poolTaobaoShopUpdateAPIResponse.Put(v)
}
