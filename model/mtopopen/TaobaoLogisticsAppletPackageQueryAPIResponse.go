package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAppletPackageQueryAPIResponse 淘宝包裹查询 API返回值
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
type TaobaoLogisticsAppletPackageQueryAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsAppletPackageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsAppletPackageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsAppletPackageQueryAPIResponseModel).Reset()
}

// TaobaoLogisticsAppletPackageQueryAPIResponseModel is 淘宝包裹查询 成功返回结果
type TaobaoLogisticsAppletPackageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_applet_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通讯失败信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果对象
	QueryResponse *QueryPackageListResponse `json:"query_response,omitempty" xml:"query_response,omitempty"`
	// 通讯成功/失败
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
	// 通讯失败码
	ResultCode bool `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsAppletPackageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.QueryResponse = nil
	m.ResultSuccess = false
	m.ResultCode = false
}

var poolTaobaoLogisticsAppletPackageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsAppletPackageQueryAPIResponse)
	},
}

// GetTaobaoLogisticsAppletPackageQueryAPIResponse 从 sync.Pool 获取 TaobaoLogisticsAppletPackageQueryAPIResponse
func GetTaobaoLogisticsAppletPackageQueryAPIResponse() *TaobaoLogisticsAppletPackageQueryAPIResponse {
	return poolTaobaoLogisticsAppletPackageQueryAPIResponse.Get().(*TaobaoLogisticsAppletPackageQueryAPIResponse)
}

// ReleaseTaobaoLogisticsAppletPackageQueryAPIResponse 将 TaobaoLogisticsAppletPackageQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsAppletPackageQueryAPIResponse(v *TaobaoLogisticsAppletPackageQueryAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsAppletPackageQueryAPIResponse.Put(v)
}
