package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
isv资源查询 APIResponse
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等
*/
type CainiaoCloudprintIsvResourcesGetAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintIsvResourcesGetResponse `json:"cainiao_cloudprint_isv_resources_get_response,omitempty"`
}

type CainiaoCloudprintIsvResourcesGetResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
