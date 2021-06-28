package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机查询订单信息 APIResponse
taobao.bus.tvmqueryorder.get

查询订单详情
*/
type TaobaoBusTvmqueryorderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmqueryorder_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"