package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单创建接口 APIResponse
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单
*/
type TaobaoQimenStoreprocessCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStoreprocessCreateResponse `json:"taobao_qimen_storeprocess_create_response,omitempty"`
}

type TaobaoQimenStoreprocessCreateResponse struct {

    // 
    Response   *StoreProcessCreateResponse `json:"response,omitempty"`

}
