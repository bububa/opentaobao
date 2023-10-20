package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiupdateAPIResponse 电子面单云打印更新接口 API返回值
// cainiao.waybill.ii.update
//
// 商家更新电子面单号对应的面单信息。
type CainiaowaybilliiupdateAPIResponse struct {
	model.CommonResponse
	CainiaowaybilliiupdateAPIResponseModel
}

// CainiaowaybilliiupdateAPIResponseModel is 电子面单云打印更新接口 成功返回结果
type CainiaowaybilliiupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 模板内容
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}
