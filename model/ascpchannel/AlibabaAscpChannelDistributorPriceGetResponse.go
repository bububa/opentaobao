package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
链渠道中心淘外分销价格查询(分销商专用) APIResponse
alibaba.ascp.channel.distributor.price.get

此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
*/
type AlibabaAscpChannelDistributorPriceGetAPIResponse struct {
    model.CommonResponse
    AlibabaAscpChannelDistributorPriceGetResponse
}

type AlibabaAscpChannelDistributorPriceGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_price_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaAscpChannelDistributorPriceGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
