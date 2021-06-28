package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
一般进口发货 APIResponse
taobao.wlb.imports.general.consign

将订单信息发送到菜鸟海外转运仓；
业务规则：
1）交易订单为待发货状态。
2）单笔订单多个商品，交易金额不能大于1000人民币。
*/
type TaobaoWlbImportsGeneralConsignAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbImportsGeneralConsignResponse `json:"wlb_imports_general_consign_response,omitempty"` 
    TaobaoWlbImportsGeneralConsignResponse
}

/* model for simplify = false
type TaobaoWlbImportsGeneralConsignResponse struct {

    // 发货成功后的物流订单编号
    
    LgorderCode   string `json:"lgorder_code,omitempty"`
    

    // 业务错误描述
    
    ResultErrorMsg   string `json:"result_error_msg,omitempty"`
    

    // 是否发货成功,true:成功，false：失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 业务错误编码
    
    ResultErrorCode   string `json:"result_error_code,omitempty"`
    

}
*/

type TaobaoWlbImportsGeneralConsignResponse struct {

    // 发货成功后的物流订单编号
    LgorderCode   string `json:"lgorder_code,omitempty"`

    // 业务错误描述
    ResultErrorMsg   string `json:"result_error_msg,omitempty"`

    // 是否发货成功,true:成功，false：失败
    IsSuccess   bool `json:"is_success,omitempty"`

    // 业务错误编码
    ResultErrorCode   string `json:"result_error_code,omitempty"`

}
