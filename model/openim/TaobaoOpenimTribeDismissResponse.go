package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群解散 APIResponse
taobao.openim.tribe.dismiss

OPENIM群解散
*/
type TaobaoOpenimTribeDismissAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeDismissResponse `json:"taobao_openim_tribe_dismiss_response,omitempty"`
}

type TaobaoOpenimTribeDismissResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
