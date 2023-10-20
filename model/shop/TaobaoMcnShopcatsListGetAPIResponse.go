package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMcnShopcatsListGetAPIResponse 店铺类目清单 API返回值
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
type TaobaoMcnShopcatsListGetAPIResponse struct {
	model.CommonResponse
	TaobaoMcnShopcatsListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMcnShopcatsListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMcnShopcatsListGetAPIResponseModel).Reset()
}

// TaobaoMcnShopcatsListGetAPIResponseModel is 店铺类目清单 成功返回结果
type TaobaoMcnShopcatsListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"mcn_shopcats_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 店铺类目列表信息
	ShopCats []ShopCat `json:"shop_cats,omitempty" xml:"shop_cats>shop_cat,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMcnShopcatsListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ShopCats = m.ShopCats[:0]
}

var poolTaobaoMcnShopcatsListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMcnShopcatsListGetAPIResponse)
	},
}

// GetTaobaoMcnShopcatsListGetAPIResponse 从 sync.Pool 获取 TaobaoMcnShopcatsListGetAPIResponse
func GetTaobaoMcnShopcatsListGetAPIResponse() *TaobaoMcnShopcatsListGetAPIResponse {
	return poolTaobaoMcnShopcatsListGetAPIResponse.Get().(*TaobaoMcnShopcatsListGetAPIResponse)
}

// ReleaseTaobaoMcnShopcatsListGetAPIResponse 将 TaobaoMcnShopcatsListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMcnShopcatsListGetAPIResponse(v *TaobaoMcnShopcatsListGetAPIResponse) {
	v.Reset()
	poolTaobaoMcnShopcatsListGetAPIResponse.Put(v)
}
