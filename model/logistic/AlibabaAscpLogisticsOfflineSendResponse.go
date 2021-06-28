package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自己联系物流发货 APIResponse
alibaba.ascp.logistics.offline.send

用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
*/
type AlibabaAscpLogisticsOfflineSendAPIResponse struct {
    model.CommonResponse
    AlibabaAscpLogisticsOfflineSendResponse
}

type AlibabaAscpLogisticsOfflineSendResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_logistics_offline_send_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaAscpLogisticsOfflineSendResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
