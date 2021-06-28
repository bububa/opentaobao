package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传一个产品，不包括产品非主图和属性图片 APIResponse
taobao.product.add

获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 <br/>传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,<br/>调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.<br/>新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
*/
type TaobaoProductAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoProductAddResponse `json:"product_add_response,omitempty"` 
    TaobaoProductAddResponse
}

/* model for simplify = false
type TaobaoProductAddResponse struct {

    // 产品结构
    
    Product  *struct {
        Product  *Product `json:"product,omitempty"`
    } `json:"product,omitempty"`
    

}
*/

type TaobaoProductAddResponse struct {

    // 产品结构
    Product   *Product `json:"product,omitempty"`

}
