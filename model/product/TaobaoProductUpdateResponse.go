package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改一个产品，可以修改主图，不能修改子图片 APIResponse
taobao.product.update

传入产品ID <br/>可修改字段：outer_id,binds,sale_props,name,price,desc,image <br/>注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG<br/>      2.商城卖家产品发布24小时后不能作删除或修改操作
*/
type TaobaoProductUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"product_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回product数据结构中的：product_id,modified
    
    Product   *Product `json:"product,omitempty" xml:"