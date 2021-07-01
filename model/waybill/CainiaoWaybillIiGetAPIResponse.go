package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单云打印接口 API返回值 
cainiao.waybill.ii.get

菜鸟电子面单的云打印申请电子面单号的方法
*/
type CainiaoWaybillIiGetAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiGetAPIResponseModel
}

// 电子面单云打印接口 成功返回结果
type CainiaoWaybillIiGetAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Modules   []WaybillCloudPrintResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_response,omitempty"`
}
