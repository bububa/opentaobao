package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductUpdateAPIResponse 修改一个产品，可以修改主图，不能修改子图片 API返回值
// taobao.product.update
//
// 传入产品ID &lt;br/&gt;可修改字段：outer_id,binds,sale_props,name,price,desc,image &lt;br/&gt;注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG&lt;br/&gt;      2.商城卖家产品发布24小时后不能作删除或修改操作
type TaobaoProductUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoProductUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductUpdateAPIResponseModel).Reset()
}

// TaobaoProductUpdateAPIResponseModel is 修改一个产品，可以修改主图，不能修改子图片 成功返回结果
type TaobaoProductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回product数据结构中的：product_id,modified
	Product *Product `json:"product,omitempty" xml:"product,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Product = nil
}

var poolTaobaoProductUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductUpdateAPIResponse)
	},
}

// GetTaobaoProductUpdateAPIResponse 从 sync.Pool 获取 TaobaoProductUpdateAPIResponse
func GetTaobaoProductUpdateAPIResponse() *TaobaoProductUpdateAPIResponse {
	return poolTaobaoProductUpdateAPIResponse.Get().(*TaobaoProductUpdateAPIResponse)
}

// ReleaseTaobaoProductUpdateAPIResponse 将 TaobaoProductUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductUpdateAPIResponse(v *TaobaoProductUpdateAPIResponse) {
	v.Reset()
	poolTaobaoProductUpdateAPIResponse.Put(v)
}
