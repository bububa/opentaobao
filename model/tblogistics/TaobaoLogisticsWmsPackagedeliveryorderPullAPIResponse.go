package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse 包裹出库单拉单 API返回值
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
type TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponseModel is 包裹出库单拉单 成功返回结果
type TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packagedeliveryorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse)
	},
}

// GetTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse
func GetTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse() *TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse {
	return poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse.Get().(*TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse)
}

// ReleaseTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse 将 TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse(v *TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse.Put(v)
}
