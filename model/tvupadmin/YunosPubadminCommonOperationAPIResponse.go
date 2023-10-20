package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosPubadminCommonOperationAPIResponse 内部迎客松通用服务 API返回值
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
type YunosPubadminCommonOperationAPIResponse struct {
	model.CommonResponse
	YunosPubadminCommonOperationAPIResponseModel
}

// Reset 清空结构体
func (m *YunosPubadminCommonOperationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosPubadminCommonOperationAPIResponseModel).Reset()
}

// YunosPubadminCommonOperationAPIResponseModel is 内部迎客松通用服务 成功返回结果
type YunosPubadminCommonOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_pubadmin_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosPubadminCommonOperationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosPubadminCommonOperationAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosPubadminCommonOperationAPIResponse)
	},
}

// GetYunosPubadminCommonOperationAPIResponse 从 sync.Pool 获取 YunosPubadminCommonOperationAPIResponse
func GetYunosPubadminCommonOperationAPIResponse() *YunosPubadminCommonOperationAPIResponse {
	return poolYunosPubadminCommonOperationAPIResponse.Get().(*YunosPubadminCommonOperationAPIResponse)
}

// ReleaseYunosPubadminCommonOperationAPIResponse 将 YunosPubadminCommonOperationAPIResponse 保存到 sync.Pool
func ReleaseYunosPubadminCommonOperationAPIResponse(v *YunosPubadminCommonOperationAPIResponse) {
	v.Reset()
	poolYunosPubadminCommonOperationAPIResponse.Put(v)
}
