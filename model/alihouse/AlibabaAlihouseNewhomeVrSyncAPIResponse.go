package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVrSyncAPIResponse VR关系数据同步 API返回值
// alibaba.alihouse.newhome.vr.sync
//
// 对接易居VR关系数据迁移
type AlibabaAlihouseNewhomeVrSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeVrSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVrSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeVrSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeVrSyncAPIResponseModel is VR关系数据同步 成功返回结果
type AlibabaAlihouseNewhomeVrSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_vr_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeVrSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVrSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeVrSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVrSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeVrSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeVrSyncAPIResponse
func GetAlibabaAlihouseNewhomeVrSyncAPIResponse() *AlibabaAlihouseNewhomeVrSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeVrSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeVrSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeVrSyncAPIResponse 将 AlibabaAlihouseNewhomeVrSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeVrSyncAPIResponse(v *AlibabaAlihouseNewhomeVrSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeVrSyncAPIResponse.Put(v)
}
