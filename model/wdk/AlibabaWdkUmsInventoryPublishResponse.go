package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
初始化覆盖实物库存 APIResponse
alibaba.wdk.ums.inventory.publish

先去库存这边查询当前实物库存有多少的量，然后算出来需要增加的量。接下来调用ums原来的入库语义的接口进行库存的增量补充
*/
type AlibabaWdkUmsInventoryPublishAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsInventoryPublishResponse
}

type AlibabaWdkUmsInventoryPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_inventory_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用服务返回结果
    
    ApiResult   *AlibabaWdkUmsInventoryPublishApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
