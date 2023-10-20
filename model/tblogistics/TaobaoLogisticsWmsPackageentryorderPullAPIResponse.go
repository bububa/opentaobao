package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageentryorderPullAPIResponse 包裹入库单拉单 API返回值
// taobao.logistics.wms.packageentryorder.pull
//
// 包裹入库单拉单
type TaobaoLogisticsWmsPackageentryorderPullAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageentryorderPullAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel is 包裹入库单拉单 成功返回结果
type TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageentryorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageentryorderPullAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsWmsPackageentryorderPullAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsPackageentryorderPullAPIResponse)
	},
}

// GetTaobaoLogisticsWmsPackageentryorderPullAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsPackageentryorderPullAPIResponse
func GetTaobaoLogisticsWmsPackageentryorderPullAPIResponse() *TaobaoLogisticsWmsPackageentryorderPullAPIResponse {
	return poolTaobaoLogisticsWmsPackageentryorderPullAPIResponse.Get().(*TaobaoLogisticsWmsPackageentryorderPullAPIResponse)
}

// ReleaseTaobaoLogisticsWmsPackageentryorderPullAPIResponse 将 TaobaoLogisticsWmsPackageentryorderPullAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageentryorderPullAPIResponse(v *TaobaoLogisticsWmsPackageentryorderPullAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageentryorderPullAPIResponse.Put(v)
}
