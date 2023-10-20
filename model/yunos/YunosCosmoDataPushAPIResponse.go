package yunos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosCosmoDataPushAPIResponse COSMO-PUSH模式数据接入 API返回值
// yunos.cosmo.data.push
//
// YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入
type YunosCosmoDataPushAPIResponse struct {
	model.CommonResponse
	YunosCosmoDataPushAPIResponseModel
}

// Reset 清空结构体
func (m *YunosCosmoDataPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosCosmoDataPushAPIResponseModel).Reset()
}

// YunosCosmoDataPushAPIResponseModel is COSMO-PUSH模式数据接入 成功返回结果
type YunosCosmoDataPushAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_cosmo_data_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DpResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosCosmoDataPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosCosmoDataPushAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosCosmoDataPushAPIResponse)
	},
}

// GetYunosCosmoDataPushAPIResponse 从 sync.Pool 获取 YunosCosmoDataPushAPIResponse
func GetYunosCosmoDataPushAPIResponse() *YunosCosmoDataPushAPIResponse {
	return poolYunosCosmoDataPushAPIResponse.Get().(*YunosCosmoDataPushAPIResponse)
}

// ReleaseYunosCosmoDataPushAPIResponse 将 YunosCosmoDataPushAPIResponse 保存到 sync.Pool
func ReleaseYunosCosmoDataPushAPIResponse(v *YunosCosmoDataPushAPIResponse) {
	v.Reset()
	poolYunosCosmoDataPushAPIResponse.Put(v)
}
