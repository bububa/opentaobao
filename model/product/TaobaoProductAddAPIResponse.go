package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductAddAPIResponse 上传一个产品，不包括产品非主图和属性图片 API返回值
// taobao.product.add
//
// 获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 &lt;br/&gt;传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,&lt;br/&gt;调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.&lt;br/&gt;新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
type TaobaoProductAddAPIResponse struct {
	model.CommonResponse
	TaobaoProductAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductAddAPIResponseModel).Reset()
}

// TaobaoProductAddAPIResponseModel is 上传一个产品，不包括产品非主图和属性图片 成功返回结果
type TaobaoProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品结构
	Product *Product `json:"product,omitempty" xml:"product,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Product = nil
}

var poolTaobaoProductAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductAddAPIResponse)
	},
}

// GetTaobaoProductAddAPIResponse 从 sync.Pool 获取 TaobaoProductAddAPIResponse
func GetTaobaoProductAddAPIResponse() *TaobaoProductAddAPIResponse {
	return poolTaobaoProductAddAPIResponse.Get().(*TaobaoProductAddAPIResponse)
}

// ReleaseTaobaoProductAddAPIResponse 将 TaobaoProductAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductAddAPIResponse(v *TaobaoProductAddAPIResponse) {
	v.Reset()
	poolTaobaoProductAddAPIResponse.Put(v)
}
