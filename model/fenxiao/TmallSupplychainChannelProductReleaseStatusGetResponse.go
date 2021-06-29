package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品铺货状态查询 APIResponse
tmall.supplychain.channel.product.release.status.get

巴拿马战役渠道产品状态查询
*/
type TmallSupplychainChannelProductReleaseStatusGetAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductReleaseStatusGetResponse
}

type TmallSupplychainChannelProductReleaseStatusGetResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_release_status_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductReleaseStatusGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
