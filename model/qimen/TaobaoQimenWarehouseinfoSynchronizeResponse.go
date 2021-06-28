package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓库同步接口 APIResponse
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
type TaobaoQimenWarehouseinfoSynchronizeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenWarehouseinfoSynchronizeResponse `json:"qimen_warehouseinfo_synchronize_response,omitempty"` 
    TaobaoQimenWarehouseinfoSynchronizeResponse
}

/* model for simplify = false
type TaobaoQimenWarehouseinfoSynchronizeResponse struct {

    // 响应报文
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenWarehouseinfoSynchronizeResponse struct {

    // 响应报文
    Response   *Response `json:"response,omitempty"`

}
