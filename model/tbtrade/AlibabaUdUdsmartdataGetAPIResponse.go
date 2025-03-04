package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaUdUdsmartdataGetAPIResponse udsmart获取数据回传接口 API返回值
// alibaba.ud.udsmartdata.get
//
// udsmart获取数据回传接口
type AlibabaUdUdsmartdataGetAPIResponse struct {
	model.CommonResponse
	AlibabaUdUdsmartdataGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaUdUdsmartdataGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaUdUdsmartdataGetAPIResponseModel).Reset()
}

// AlibabaUdUdsmartdataGetAPIResponseModel is udsmart获取数据回传接口 成功返回结果
type AlibabaUdUdsmartdataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ud_udsmartdata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// request_id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// CallBackResultDTO
	Data *CallBackResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaUdUdsmartdataGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Id = ""
	m.Data = nil
	m.Status = 0
}

var poolAlibabaUdUdsmartdataGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaUdUdsmartdataGetAPIResponse)
	},
}

// GetAlibabaUdUdsmartdataGetAPIResponse 从 sync.Pool 获取 AlibabaUdUdsmartdataGetAPIResponse
func GetAlibabaUdUdsmartdataGetAPIResponse() *AlibabaUdUdsmartdataGetAPIResponse {
	return poolAlibabaUdUdsmartdataGetAPIResponse.Get().(*AlibabaUdUdsmartdataGetAPIResponse)
}

// ReleaseAlibabaUdUdsmartdataGetAPIResponse 将 AlibabaUdUdsmartdataGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaUdUdsmartdataGetAPIResponse(v *AlibabaUdUdsmartdataGetAPIResponse) {
	v.Reset()
	poolAlibabaUdUdsmartdataGetAPIResponse.Put(v)
}
