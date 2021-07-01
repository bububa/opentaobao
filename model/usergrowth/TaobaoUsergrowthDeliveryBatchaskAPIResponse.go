package usergrowth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放批量询问 API返回值 
taobao.usergrowth.delivery.batchask

提供给媒体在曝光广告前调用， 返回是否曝光以及报价
*/
type TaobaoUsergrowthDeliveryBatchaskAPIResponse struct {
    model.CommonResponse
    TaobaoUsergrowthDeliveryBatchaskAPIResponseModel
}

// 广告投放批量询问 成功返回结果
type TaobaoUsergrowthDeliveryBatchaskAPIResponseModel struct {
    XMLName xml.Name `xml:"usergrowth_delivery_batchask_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值， 每个目标用户对应一个接口，具体参见示例
    Results   []AskResult `json:"results,omitempty" xml:"results>ask_result,omitempty"`
}
