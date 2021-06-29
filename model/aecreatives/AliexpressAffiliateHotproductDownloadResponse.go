package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟营销爆品下载接口 APIResponse
aliexpress.affiliate.hotproduct.download

查询联盟爆品API
*/
type AliexpressAffiliateHotproductDownloadAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateHotproductDownloadResponse
}

type AliexpressAffiliateHotproductDownloadResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_download_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
