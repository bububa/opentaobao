package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDatabankOpenOneserviceGetdataAPIResponse 瓴羊DaaS消费者运营CGP取数接口 API返回值
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
type AlibabaDatabankOpenOneserviceGetdataAPIResponse struct {
	model.CommonResponse
	AlibabaDatabankOpenOneserviceGetdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDatabankOpenOneserviceGetdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDatabankOpenOneserviceGetdataAPIResponseModel).Reset()
}

// AlibabaDatabankOpenOneserviceGetdataAPIResponseModel is 瓴羊DaaS消费者运营CGP取数接口 成功返回结果
type AlibabaDatabankOpenOneserviceGetdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_databank_open_oneservice_getdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 指标结果
	Data []string `json:"data,omitempty" xml:"data>string,omitempty"`
	// 空
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 请求成功
	CodeClass string `json:"code_class,omitempty" xml:"code_class,omitempty"`
	// 请求成功
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDatabankOpenOneserviceGetdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.Errmsg = ""
	m.CodeClass = ""
	m.Errcode = 0
}

var poolAlibabaDatabankOpenOneserviceGetdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDatabankOpenOneserviceGetdataAPIResponse)
	},
}

// GetAlibabaDatabankOpenOneserviceGetdataAPIResponse 从 sync.Pool 获取 AlibabaDatabankOpenOneserviceGetdataAPIResponse
func GetAlibabaDatabankOpenOneserviceGetdataAPIResponse() *AlibabaDatabankOpenOneserviceGetdataAPIResponse {
	return poolAlibabaDatabankOpenOneserviceGetdataAPIResponse.Get().(*AlibabaDatabankOpenOneserviceGetdataAPIResponse)
}

// ReleaseAlibabaDatabankOpenOneserviceGetdataAPIResponse 将 AlibabaDatabankOpenOneserviceGetdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDatabankOpenOneserviceGetdataAPIResponse(v *AlibabaDatabankOpenOneserviceGetdataAPIResponse) {
	v.Reset()
	poolAlibabaDatabankOpenOneserviceGetdataAPIResponse.Put(v)
}
