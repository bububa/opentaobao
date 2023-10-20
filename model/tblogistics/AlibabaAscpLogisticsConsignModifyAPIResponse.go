package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsConsignModifyAPIResponse 修改物流公司和运单号 API返回值
// alibaba.ascp.logistics.consign.modify
//
// 修改物流公司和运单号
type AlibabaAscpLogisticsConsignModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsConsignModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsConsignModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsConsignModifyAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsConsignModifyAPIResponseModel is 修改物流公司和运单号 成功返回结果
type AlibabaAscpLogisticsConsignModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_consign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsConsignModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsConsignModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsConsignModifyAPIResponse)
	},
}

// GetAlibabaAscpLogisticsConsignModifyAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsConsignModifyAPIResponse
func GetAlibabaAscpLogisticsConsignModifyAPIResponse() *AlibabaAscpLogisticsConsignModifyAPIResponse {
	return poolAlibabaAscpLogisticsConsignModifyAPIResponse.Get().(*AlibabaAscpLogisticsConsignModifyAPIResponse)
}

// ReleaseAlibabaAscpLogisticsConsignModifyAPIResponse 将 AlibabaAscpLogisticsConsignModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsConsignModifyAPIResponse(v *AlibabaAscpLogisticsConsignModifyAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsConsignModifyAPIResponse.Put(v)
}
