package cainiaolocker

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoWaybillGetAPIResponse 菜鸟裹裹商家寄件取号接口 API返回值
// cainiao.guoguo.waybill.get
//
// 菜鸟裹裹商家寄件取号接口
type CainiaoGuoguoWaybillGetAPIResponse struct {
	model.CommonResponse
	CainiaoGuoguoWaybillGetAPIResponseModel
}

// CainiaoGuoguoWaybillGetAPIResponseModel is 菜鸟裹裹商家寄件取号接口 成功返回结果
type CainiaoGuoguoWaybillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_waybill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Modules []WaybillCloudPrintNewResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_new_response,omitempty"`
}
