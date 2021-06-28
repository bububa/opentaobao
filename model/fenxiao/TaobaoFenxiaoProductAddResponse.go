package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加产品 APIResponse
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
*/
type TaobaoFenxiaoProductAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductAddResponse `json:"fenxiao_product_add_response,omitempty"` 
    TaobaoFenxiaoProductAddResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductAddResponse struct {

    // 产品ID
    
    Pid   int64 `json:"pid,omitempty"`
    

    // 产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss
    
    Created   string `json:"created,omitempty"`
    

}
*/

type TaobaoFenxiaoProductAddResponse struct {

    // 产品ID
    Pid   int64 `json:"pid,omitempty"`

    // 产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss
    Created   string `json:"created,omitempty"`

}
