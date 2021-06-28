package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的自定义区模板信息 APIResponse
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息
*/
type CainiaoCloudprintCustomaresGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_customares_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"