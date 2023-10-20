package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoServicePackageQueryAPIResponse 预采购服务包查询接口 API返回值
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
type AlibabaEleFengniaoServicePackageQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoServicePackageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoServicePackageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoServicePackageQueryAPIResponseModel).Reset()
}

// AlibabaEleFengniaoServicePackageQueryAPIResponseModel is 预采购服务包查询接口 成功返回结果
type AlibabaEleFengniaoServicePackageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_service_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// servicePackages
	ServicePackages []AlibabaEleFengniaoServicePackageQueryResult `json:"service_packages,omitempty" xml:"service_packages>alibaba_ele_fengniao_service_package_query_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoServicePackageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServicePackages = m.ServicePackages[:0]
}

var poolAlibabaEleFengniaoServicePackageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoServicePackageQueryAPIResponse)
	},
}

// GetAlibabaEleFengniaoServicePackageQueryAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoServicePackageQueryAPIResponse
func GetAlibabaEleFengniaoServicePackageQueryAPIResponse() *AlibabaEleFengniaoServicePackageQueryAPIResponse {
	return poolAlibabaEleFengniaoServicePackageQueryAPIResponse.Get().(*AlibabaEleFengniaoServicePackageQueryAPIResponse)
}

// ReleaseAlibabaEleFengniaoServicePackageQueryAPIResponse 将 AlibabaEleFengniaoServicePackageQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoServicePackageQueryAPIResponse(v *AlibabaEleFengniaoServicePackageQueryAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoServicePackageQueryAPIResponse.Put(v)
}
