package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSkusCustomGetAPIResponse 根据外部ID取商品SKU API返回值
// taobao.skus.custom.get
//
// 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku <br/>这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoSkusCustomGetAPIResponse struct {
	model.CommonResponse
	TaobaoSkusCustomGetAPIResponseModel
}

// TaobaoSkusCustomGetAPIResponseModel is 根据外部ID取商品SKU 成功返回结果
type TaobaoSkusCustomGetAPIResponseModel struct {
	XMLName xml.Name `xml:"skus_custom_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Sku对象，具体字段以fields决定
	Skus []Sku `json:"skus,omitempty" xml:"skus>sku,omitempty"`
}
