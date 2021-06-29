package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘外分销退货回传物流单号 APIResponse
alibaba.ascp.channel.refund.goods.waybill

淘外分销退货回传物流单号
*/
type AlibabaAscpChannelRefundGoodsWaybillAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelRefundGoodsWaybillResponse
}

type AlibabaAscpChannelRefundGoodsWaybillResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_refund_goods_waybill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaAscpChannelRefundGoodsWaybillResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
