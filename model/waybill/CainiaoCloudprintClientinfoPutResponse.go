package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印客户端监控信息收集 APIResponse
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集
*/
type CainiaoCloudprintClientinfoPutAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_clientinfo_put_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"