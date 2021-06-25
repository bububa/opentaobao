package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
异常订单信息获取 APIResponse
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口
*/
type TaobaoOcTradetraceAlertsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOcTradetraceAlertsGetResponse `json:"taobao_oc_tradetrace_alerts_get_response,omitempty"`
}

type TaobaoOcTradetraceAlertsGetResponse struct {

    // 异常订单数据
    ResultList   []SimpleAbnormalOrderDetail `json:"result_list,omitempty"`

}
