package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsSendinfoQueryAPIResponse 供应商-异云-查询主订单包含的物流单 API返回值
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
type AlibabaLstLogisticsSendinfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstLogisticsSendinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsSendinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstLogisticsSendinfoQueryAPIResponseModel).Reset()
}

// AlibabaLstLogisticsSendinfoQueryAPIResponseModel is 供应商-异云-查询主订单包含的物流单 成功返回结果
type AlibabaLstLogisticsSendinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_sendinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstLogisticsSendinfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstLogisticsSendinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstLogisticsSendinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstLogisticsSendinfoQueryAPIResponse)
	},
}

// GetAlibabaLstLogisticsSendinfoQueryAPIResponse 从 sync.Pool 获取 AlibabaLstLogisticsSendinfoQueryAPIResponse
func GetAlibabaLstLogisticsSendinfoQueryAPIResponse() *AlibabaLstLogisticsSendinfoQueryAPIResponse {
	return poolAlibabaLstLogisticsSendinfoQueryAPIResponse.Get().(*AlibabaLstLogisticsSendinfoQueryAPIResponse)
}

// ReleaseAlibabaLstLogisticsSendinfoQueryAPIResponse 将 AlibabaLstLogisticsSendinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstLogisticsSendinfoQueryAPIResponse(v *AlibabaLstLogisticsSendinfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstLogisticsSendinfoQueryAPIResponse.Put(v)
}
