package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品 APIResponse
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品
*/
type TaobaoBanamadpcItemUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBanamadpcItemUpdateResponse `json:"taobao_banamadpc_item_update_response,omitempty"`
}

type TaobaoBanamadpcItemUpdateResponse struct {

    // 无
    ApiResult   *TaobaoBanamadpcItemUpdateApiResult `json:"api_result,omitempty"`

}
