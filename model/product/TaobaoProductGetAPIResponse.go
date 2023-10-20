package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductGetAPIResponse 获取一个产品的信息 API返回值
// taobao.product.get
//
// 天猫商家发布商品时，查询关联产品信息时使用，非商品查询接口。商品查询接口：taobao.item.seller.get&lt;br&gt;
// 两种方式查看一个产品详细信息:
// 传入product_id来查询；传入cid和props来查询
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoProductGetAPIResponse struct {
	model.CommonResponse
	TaobaoProductGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductGetAPIResponseModel).Reset()
}

// TaobaoProductGetAPIResponseModel is 获取一个产品的信息 成功返回结果
type TaobaoProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回具体信息为入参fields请求的字段信息
	Product *Product `json:"product,omitempty" xml:"product,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Product = nil
}

var poolTaobaoProductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductGetAPIResponse)
	},
}

// GetTaobaoProductGetAPIResponse 从 sync.Pool 获取 TaobaoProductGetAPIResponse
func GetTaobaoProductGetAPIResponse() *TaobaoProductGetAPIResponse {
	return poolTaobaoProductGetAPIResponse.Get().(*TaobaoProductGetAPIResponse)
}

// ReleaseTaobaoProductGetAPIResponse 将 TaobaoProductGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductGetAPIResponse(v *TaobaoProductGetAPIResponse) {
	v.Reset()
	poolTaobaoProductGetAPIResponse.Put(v)
}
