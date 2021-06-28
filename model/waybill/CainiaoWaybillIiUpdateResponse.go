package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印更新接口 APIResponse
cainiao.waybill.ii.update

商家更新电子面单号对应的面单信息。
*/
type CainiaoWaybillIiUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_ii_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 面单号
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"