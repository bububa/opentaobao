package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单确认接口 APIResponse
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况
*/
type TaobaoQimenStoreprocessConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStoreprocessConfirmResponse `json:"taobao_qimen_storeprocess_confirm_response,omitempty"`
}

type TaobaoQimenStoreprocessConfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
