package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告曝光前判定批量接口V2 APIResponse
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2
*/
type TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDhhDeliveryBatchaskResponse
}

type TaobaoUsergrowthDhhDeliveryBatchaskResponse struct {
    XMLName xml.Name `xml:"usergrowth_dhh_delivery_batchask_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *BatchAskResultV2 `json:"result,omitempty" xml:"result,omitempty"`

    
}
