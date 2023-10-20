package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercatsListAddAPIResponse 添加卖家自定义类目 API返回值
// taobao.sellercats.list.add
//
// 此API添加卖家店铺内自定义类目 &lt;br/&gt;父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 &lt;br/&gt;注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
type TaobaoSellercatsListAddAPIResponse struct {
	model.CommonResponse
	TaobaoSellercatsListAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSellercatsListAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSellercatsListAddAPIResponseModel).Reset()
}

// TaobaoSellercatsListAddAPIResponseModel is 添加卖家自定义类目 成功返回结果
type TaobaoSellercatsListAddAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercats_list_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回seller_cat数据结构中的：cid,created
	SellerCat *SellerCat `json:"seller_cat,omitempty" xml:"seller_cat,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSellercatsListAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SellerCat = nil
}

var poolTaobaoSellercatsListAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSellercatsListAddAPIResponse)
	},
}

// GetTaobaoSellercatsListAddAPIResponse 从 sync.Pool 获取 TaobaoSellercatsListAddAPIResponse
func GetTaobaoSellercatsListAddAPIResponse() *TaobaoSellercatsListAddAPIResponse {
	return poolTaobaoSellercatsListAddAPIResponse.Get().(*TaobaoSellercatsListAddAPIResponse)
}

// ReleaseTaobaoSellercatsListAddAPIResponse 将 TaobaoSellercatsListAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSellercatsListAddAPIResponse(v *TaobaoSellercatsListAddAPIResponse) {
	v.Reset()
	poolTaobaoSellercatsListAddAPIResponse.Put(v)
}
