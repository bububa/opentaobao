package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店更新接口 APIResponse
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息
*/
type TaobaoQimenStoreUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenStoreUpdateResponse `json:"taobao_qimen_store_update_response,omitempty"`
}

type TaobaoQimenStoreUpdateResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应标示
    Flag   string `json:"flag,omitempty"`

    // 响应编码
    QimenCode   string `json:"qimen_code,omitempty"`

}
