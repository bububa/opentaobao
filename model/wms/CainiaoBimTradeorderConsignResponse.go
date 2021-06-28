package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
驱动保税交易订单发货 APIResponse
cainiao.bim.tradeorder.consign

驱动保税交易订单发货
*/
type CainiaoBimTradeorderConsignAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoBimTradeorderConsignResponse `json:"cainiao_bim_tradeorder_consign_response,omitempty"` 
    CainiaoBimTradeorderConsignResponse
}

/* model for simplify = false
type CainiaoBimTradeorderConsignResponse struct {

    // 菜鸟仓库作业单据号
    
    StoreOrderCode   string `json:"store_order_code,omitempty"`
    

    // 菜鸟物流订单号,格式为LP开头
    
    LgOrderCode   string `json:"lg_order_code,omitempty"`
    

}
*/

type CainiaoBimTradeorderConsignResponse struct {

    // 菜鸟仓库作业单据号
    StoreOrderCode   string `json:"store_order_code,omitempty"`

    // 菜鸟物流订单号,格式为LP开头
    LgOrderCode   string `json:"lg_order_code,omitempty"`

}
