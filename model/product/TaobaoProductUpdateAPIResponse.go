package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoproductupdateAPIResponse 修改一个产品，可以修改主图，不能修改子图片 API返回值
// taobao.product.update
//
// 传入产品ID &lt;br/&gt;可修改字段：outer_id,binds,sale_props,name,price,desc,image &lt;br/&gt;注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG&lt;br/&gt;      2.商城卖家产品发布24小时后不能作删除或修改操作
type TaobaoproductupdateAPIResponse struct {
	model.CommonResponse
	TaobaoproductupdateAPIResponseModel
}

// TaobaoproductupdateAPIResponseModel is 修改一个产品，可以修改主图，不能修改子图片 成功返回结果
type TaobaoproductupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回product数据结构中的：product_id,modified
	Product *Product `json:"product,omitempty" xml:"product,omitempty"`
}
