package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单明细上传接口 APIResponse
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据
*/
type TaobaoFuwuSpBillreordAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFuwuSpBillreordAddResponse `json:"fuwu_sp_billreord_add_response,omitempty"` 
    TaobaoFuwuSpBillreordAddResponse
}

/* model for simplify = false
type TaobaoFuwuSpBillreordAddResponse struct {

    // 返回调用结果
    
    AddResult   bool `json:"add_result,omitempty"`
    

}
*/

type TaobaoFuwuSpBillreordAddResponse struct {

    // 返回调用结果
    AddResult   bool `json:"add_result,omitempty"`

}
