package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有的菜鸟标准电子面单模板 APIResponse
cainiao.cloudprint.stdtemplates.get

获取菜鸟标准电子面单模板
*/
type CainiaoCloudprintStdtemplatesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_stdtemplates_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果集
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"