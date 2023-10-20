package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTobPackageQueryAPIResponse B2B包裹查询接口 API返回值
// alibaba.ascp.uop.tob.package.query
//
// 供应链中台TOB包裹查询接口
type AlibabaAscpUopTobPackageQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopTobPackageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopTobPackageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopTobPackageQueryAPIResponseModel).Reset()
}

// AlibabaAscpUopTobPackageQueryAPIResponseModel is B2B包裹查询接口 成功返回结果
type AlibabaAscpUopTobPackageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_tob_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	PackageQueryResponse *ResultWrapper `json:"package_query_response,omitempty" xml:"package_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopTobPackageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PackageQueryResponse = nil
}

var poolAlibabaAscpUopTobPackageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopTobPackageQueryAPIResponse)
	},
}

// GetAlibabaAscpUopTobPackageQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpUopTobPackageQueryAPIResponse
func GetAlibabaAscpUopTobPackageQueryAPIResponse() *AlibabaAscpUopTobPackageQueryAPIResponse {
	return poolAlibabaAscpUopTobPackageQueryAPIResponse.Get().(*AlibabaAscpUopTobPackageQueryAPIResponse)
}

// ReleaseAlibabaAscpUopTobPackageQueryAPIResponse 将 AlibabaAscpUopTobPackageQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopTobPackageQueryAPIResponse(v *AlibabaAscpUopTobPackageQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopTobPackageQueryAPIResponse.Put(v)
}
