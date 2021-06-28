package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
异常订单信息获取 APIResponse
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口
*/
type TaobaoOcTradetraceAlertsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_tradetrace_alerts_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异常订单数据
    
    ResultList   []SimpleAbnormalOrderDetail `json:"result_list,omitempty" xml:"