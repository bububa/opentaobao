package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云打印客户端监控信息收集 APIResponse
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集
*/
type CainiaoCloudprintClientinfoPutAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintClientinfoPutResponse `json:"cainiao_cloudprint_clientinfo_put_response,omitempty"`
}

type CainiaoCloudprintClientinfoPutResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
