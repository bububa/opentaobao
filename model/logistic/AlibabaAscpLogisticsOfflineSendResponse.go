package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
自己联系物流发货 APIResponse
alibaba.ascp.logistics.offline.send

用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
*/
type AlibabaAscpLogisticsOfflineSendAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAscpLogisticsOfflineSendResponse `json:"alibaba_ascp_logistics_offline_send_response,omitempty"`
}

type AlibabaAscpLogisticsOfflineSendResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlibabaAscpLogisticsOfflineSendResultDto `json:"result,omitempty"`

}
