package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印接口 APIResponse
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
type CainiaoWaybillIiGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_ii_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Modules   []WaybillCloudPrintResponse `json:"modules,omitempty" xml:"