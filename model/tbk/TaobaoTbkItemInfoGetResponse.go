package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-淘宝客商品详情查询(简版) APIResponse
taobao.tbk.item.info.get

淘宝客商品详情查询（简版）
*/
type TaobaoTbkItemInfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkItemInfoGetResponse `json:"taobao_tbk_item_info_get_response,omitempty"`
}

type TaobaoTbkItemInfoGetResponse struct {

    // 淘宝客商品
    Results   []NTbkItem `json:"results,omitempty"`

}
