package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryIcpQueryLbxAPIResponse icp订单号查询lbx订单号 API返回值
// alibaba.ascp.industry.icp.query.lbx
//
// 根据icp订单号查询lbx订单号
type AlibabaAscpIndustryIcpQueryLbxAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryIcpQueryLbxAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryIcpQueryLbxAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryIcpQueryLbxAPIResponseModel).Reset()
}

// AlibabaAscpIndustryIcpQueryLbxAPIResponseModel is icp订单号查询lbx订单号 成功返回结果
type AlibabaAscpIndustryIcpQueryLbxAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_icp_query_lbx_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	BizResponse *ResultWrapper `json:"biz_response,omitempty" xml:"biz_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryIcpQueryLbxAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizResponse = nil
}

var poolAlibabaAscpIndustryIcpQueryLbxAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryIcpQueryLbxAPIResponse)
	},
}

// GetAlibabaAscpIndustryIcpQueryLbxAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryIcpQueryLbxAPIResponse
func GetAlibabaAscpIndustryIcpQueryLbxAPIResponse() *AlibabaAscpIndustryIcpQueryLbxAPIResponse {
	return poolAlibabaAscpIndustryIcpQueryLbxAPIResponse.Get().(*AlibabaAscpIndustryIcpQueryLbxAPIResponse)
}

// ReleaseAlibabaAscpIndustryIcpQueryLbxAPIResponse 将 AlibabaAscpIndustryIcpQueryLbxAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryIcpQueryLbxAPIResponse(v *AlibabaAscpIndustryIcpQueryLbxAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryIcpQueryLbxAPIResponse.Put(v)
}
