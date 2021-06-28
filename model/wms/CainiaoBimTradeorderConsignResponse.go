package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
驱动保税交易订单发货 APIResponse
cainiao.bim.tradeorder.consign

驱动保税交易订单发货
*/
type CainiaoBimTradeorderConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_bim_tradeorder_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 菜鸟仓库作业单据号
    
    StoreOrderCode   string `json:"store_order_code,omitempty" xml:"