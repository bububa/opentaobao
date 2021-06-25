package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取授权账号对应的OpenUid APIResponse
taobao.openuid.get

获取授权账号对应的OpenUid
*/
type TaobaoOpenuidGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenuidGetResponse `json:"taobao_openuid_get_response,omitempty"`
}

type TaobaoOpenuidGetResponse struct {

    // OpenUID
    OpenUid   string `json:"open_uid,omitempty"`

}
