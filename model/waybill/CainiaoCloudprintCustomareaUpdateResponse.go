package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
自定义区内容更新 APIResponse
cainiao.cloudprint.customarea.update

自定义区内容更新
*/
type CainiaoCloudprintCustomareaUpdateAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCloudprintCustomareaUpdateResponse `json:"cainiao_cloudprint_customarea_update_response,omitempty"`
}

type CainiaoCloudprintCustomareaUpdateResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
