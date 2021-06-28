package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川支付完成回调 APIResponse
taobao.baichuan.payresult.query

百川支付完成回调
*/
type TaobaoBaichuanPayresultQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanPayresultQueryResponse `json:"baichuan_payresult_query_response,omitempty"` 
    TaobaoBaichuanPayresultQueryResponse
}

/* model for simplify = false
type TaobaoBaichuanPayresultQueryResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanPayresultQueryResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
