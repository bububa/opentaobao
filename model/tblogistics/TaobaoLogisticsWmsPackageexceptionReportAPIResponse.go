package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackageexceptionReportAPIResponse 无主件回告 API返回值
// taobao.logistics.wms.packageexception.report
//
// 无主件回告
type TaobaoLogisticsWmsPackageexceptionReportAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsPackageexceptionReportAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageexceptionReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsPackageexceptionReportAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsPackageexceptionReportAPIResponseModel is 无主件回告 成功返回结果
type TaobaoLogisticsWmsPackageexceptionReportAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_packageexception_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsPackageexceptionReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsWmsPackageexceptionReportAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsPackageexceptionReportAPIResponse)
	},
}

// GetTaobaoLogisticsWmsPackageexceptionReportAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsPackageexceptionReportAPIResponse
func GetTaobaoLogisticsWmsPackageexceptionReportAPIResponse() *TaobaoLogisticsWmsPackageexceptionReportAPIResponse {
	return poolTaobaoLogisticsWmsPackageexceptionReportAPIResponse.Get().(*TaobaoLogisticsWmsPackageexceptionReportAPIResponse)
}

// ReleaseTaobaoLogisticsWmsPackageexceptionReportAPIResponse 将 TaobaoLogisticsWmsPackageexceptionReportAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsPackageexceptionReportAPIResponse(v *TaobaoLogisticsWmsPackageexceptionReportAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsPackageexceptionReportAPIResponse.Put(v)
}
