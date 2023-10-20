package tblogistics

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel is 包裹入库单确认 成功返回结果
type TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageentryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrappe `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageentryorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse)
	},
}

// GetTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse
func GetTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse() *TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse {
	return poolTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse.Get().(*TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse)
}

// ReleaseTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse 将 TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse(v *TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageentryorderConfirmAPIResponse.Put(v)
}
