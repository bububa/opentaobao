package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID或商家编码修改后端商品 APIResponse
taobao.scitem.update

根据商品ID或商家编码修改后端商品
*/
type TaobaoScitemUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemUpdateResponse `json:"scitem_update_response,omitempty"` 
    TaobaoScitemUpdateResponse
}

/* model for simplify = false
type TaobaoScitemUpdateResponse struct {

    // 更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据
    
    UpdateRows   int64 `json:"update_rows,omitempty"`
    

}
*/

type TaobaoScitemUpdateResponse struct {

    // 更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据
    UpdateRows   int64 `json:"update_rows,omitempty"`

}
