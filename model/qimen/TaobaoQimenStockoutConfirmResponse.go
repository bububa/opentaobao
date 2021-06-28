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
    TaobaoQimenStockoutConfirmResponse
}

type TaobaoQimenStockoutConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_stockout_confirm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenStockoutConfirmStruct `json:"response,omitempty" xml:"response,omitempty"`

    
}
