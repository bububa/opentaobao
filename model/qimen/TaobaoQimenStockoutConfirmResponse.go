package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单确认接口 APIResponse
taobao.qimen.stockout.confirm

货品出库后，WMS将状态回传给ERP
*/
type TaobaoQimenStockoutConfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_stockout_confirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenStockoutConfirmStruct `json:"response,omitempty" xml:"