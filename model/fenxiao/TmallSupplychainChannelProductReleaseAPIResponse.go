package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商铺货 API返回值 
tmall.supplychain.channel.product.release

供应商渠道铺货接口
*/
type TmallSupplychainChannelProductReleaseAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductReleaseAPIResponseModel
}

// 供应商铺货 成功返回结果
type TmallSupplychainChannelProductReleaseAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_release_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TmallSupplychainChannelProductReleaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
