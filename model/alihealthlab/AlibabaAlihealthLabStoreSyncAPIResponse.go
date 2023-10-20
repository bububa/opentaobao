package alihealthlab

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabStoreSyncAPIResponse 阿里健康检验检测业务，isv门店同步到健康 API返回值
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
type AlibabaAlihealthLabStoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthLabStoreSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabStoreSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthLabStoreSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthLabStoreSyncAPIResponseModel is 阿里健康检验检测业务，isv门店同步到健康 成功返回结果
type AlibabaAlihealthLabStoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_lab_store_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 可读的错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthLabStoreSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultStatus = ""
	m.ResultCode = ""
}

var poolAlibabaAlihealthLabStoreSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthLabStoreSyncAPIResponse)
	},
}

// GetAlibabaAlihealthLabStoreSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthLabStoreSyncAPIResponse
func GetAlibabaAlihealthLabStoreSyncAPIResponse() *AlibabaAlihealthLabStoreSyncAPIResponse {
	return poolAlibabaAlihealthLabStoreSyncAPIResponse.Get().(*AlibabaAlihealthLabStoreSyncAPIResponse)
}

// ReleaseAlibabaAlihealthLabStoreSyncAPIResponse 将 AlibabaAlihealthLabStoreSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthLabStoreSyncAPIResponse(v *AlibabaAlihealthLabStoreSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthLabStoreSyncAPIResponse.Put(v)
}
