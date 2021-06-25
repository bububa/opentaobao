package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印接口 APIResponse
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
type CainiaoWaybillIiGetAPIResponse struct {
    model.CommonResponse
    Response *CainiaoWaybillIiGetResponse `json:"cainiao_waybill_ii_get_response,omitempty"`
}

type CainiaoWaybillIiGetResponse struct {

    // 系统自动生成
    Modules   []WaybillCloudPrintResponse `json:"modules,omitempty"`

}
