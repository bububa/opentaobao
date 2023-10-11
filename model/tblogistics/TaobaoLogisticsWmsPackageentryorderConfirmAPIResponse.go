package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse 包裹入库单确认 API返回值
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
type TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel
}

// TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel is 包裹入库单确认 成功返回结果
type TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageentryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrappe `json:"result,omitempty" xml:"result,omitempty"`
}
