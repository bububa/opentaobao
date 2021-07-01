package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductAddAPIResponse
上传一个产品，不包括产品非主图和属性图片 API返回值
taobao.product.add

获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 <br/>传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,<br/>调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.<br/>新增：套装产品发布，目前支持单件多个即 A*2 形式的套装 */
type TaobaoProductAddAPIResponse struct {
	model.CommonResponse
	TaobaoProductAddAPIResponseModel
}

// TaobaoProductAddAPIResponseModel is 上传一个产品，不包括产品非主图和属性图片 成功返回结果
type TaobaoProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品结构
	Product *Product `json:"product,omitempty" xml:"product,omitempty"`
}
