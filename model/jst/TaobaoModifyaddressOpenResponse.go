package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝自助修改地址服务开通 APIResponse
taobao.modifyaddress.open

商家自助修改地址服务开通
*/
type TaobaoModifyaddressOpenAPIResponse struct {
    model.CommonResponse
    Response *TaobaoModifyaddressOpenResponse `json:"taobao_modifyaddress_open_response,omitempty"`
}

type TaobaoModifyaddressOpenResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

}
