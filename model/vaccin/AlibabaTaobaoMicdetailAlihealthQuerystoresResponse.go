package vaccin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约门店列表查询 APIResponse
alibaba.taobao.micdetail.alihealth.querystores

微信小程序疫苗预约门店列表查询
*/
type AlibabaTaobaoMicdetailAlihealthQuerystoresAPIResponse struct {
    model.CommonResponse
    AlibabaTaobaoMicdetailAlihealthQuerystoresResponse
}

type AlibabaTaobaoMicdetailAlihealthQuerystoresResponse struct {
    XMLName xml.Name `xml:"alibaba_taobao_micdetail_alihealth_querystores_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaTaobaoMicdetailAlihealthQuerystoresResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
