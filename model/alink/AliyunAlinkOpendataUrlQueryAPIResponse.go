package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAlinkOpendataUrlQueryAPIResponse 开放数据授权访问URL查询 API返回值
// aliyun.alink.opendata.url.query
//
// 厂商数据授权访问URL查询
type AliyunAlinkOpendataUrlQueryAPIResponse struct {
	model.CommonResponse
	AliyunAlinkOpendataUrlQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAlinkOpendataUrlQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAlinkOpendataUrlQueryAPIResponseModel).Reset()
}

// AliyunAlinkOpendataUrlQueryAPIResponseModel is 开放数据授权访问URL查询 成功返回结果
type AliyunAlinkOpendataUrlQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alink_opendata_url_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权url
	Module []string `json:"module,omitempty" xml:"module>string,omitempty"`
	// 接口描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAlinkOpendataUrlQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = m.Module[:0]
	m.Message = ""
	m.Status = 0
	m.IsSuccess = false
}

var poolAliyunAlinkOpendataUrlQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAlinkOpendataUrlQueryAPIResponse)
	},
}

// GetAliyunAlinkOpendataUrlQueryAPIResponse 从 sync.Pool 获取 AliyunAlinkOpendataUrlQueryAPIResponse
func GetAliyunAlinkOpendataUrlQueryAPIResponse() *AliyunAlinkOpendataUrlQueryAPIResponse {
	return poolAliyunAlinkOpendataUrlQueryAPIResponse.Get().(*AliyunAlinkOpendataUrlQueryAPIResponse)
}

// ReleaseAliyunAlinkOpendataUrlQueryAPIResponse 将 AliyunAlinkOpendataUrlQueryAPIResponse 保存到 sync.Pool
func ReleaseAliyunAlinkOpendataUrlQueryAPIResponse(v *AliyunAlinkOpendataUrlQueryAPIResponse) {
	v.Reset()
	poolAliyunAlinkOpendataUrlQueryAPIResponse.Put(v)
}
