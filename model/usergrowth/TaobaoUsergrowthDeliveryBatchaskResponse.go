package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放批量询问 APIResponse
taobao.usergrowth.delivery.batchask

提供给媒体在曝光广告前调用， 返回是否曝光以及报价
*/
type TaobaoUsergrowthDeliveryBatchaskAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDeliveryBatchaskResponse
}

type TaobaoUsergrowthDeliveryBatchaskResponse struct {
    XMLName xml.Name `xml:"usergrowth_delivery_batchask_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值， 每个目标用户对应一个接口，具体参见示例
    
    Results   []AskResult `json:"results,omitempty" xml:"results>ask_result,omitempty"`
    
    
}
