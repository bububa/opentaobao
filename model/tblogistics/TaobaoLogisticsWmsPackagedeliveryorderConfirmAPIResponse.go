package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse 包裹出库单确认 API返回值
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
type TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponseModel is 包裹出库单确认 成功返回结果
type TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packagedeliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse)
	},
}

// GetTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse
func GetTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse() *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse {
	return poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse.Get().(*TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse)
}

// ReleaseTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse 将 TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse(v *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse.Put(v)
}
