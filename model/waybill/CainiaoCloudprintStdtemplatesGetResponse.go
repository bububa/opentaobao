package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取所有的菜鸟标准电子面单模板 APIResponse
cainiao.cloudprint.stdtemplates.get

获取菜鸟标准电子面单模板
*/
type CainiaoCloudprintStdtemplatesGetAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintStdtemplatesGetResponse `json:"cainiao_cloudprint_stdtemplates_get_response,omitempty"`
}

type CainiaoCloudprintStdtemplatesGetResponse struct {

    // 结果集
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
