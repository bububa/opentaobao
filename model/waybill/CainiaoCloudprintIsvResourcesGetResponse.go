package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
isv资源查询 APIResponse
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等
*/
type CainiaoCloudprintIsvResourcesGetAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintIsvResourcesGetResponse
}

type CainiaoCloudprintIsvResourcesGetResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_isv_resources_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
