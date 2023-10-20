package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderAllocatedinfoSyncAPIResponse 分单结果同步给星盘 API返回值
// taobao.omniorder.allocatedinfo.sync
//
// ISV分单完成，将分单结果同步给星盘
type TaobaoOmniorderAllocatedinfoSyncAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderAllocatedinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderAllocatedinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderAllocatedinfoSyncAPIResponseModel).Reset()
}

// TaobaoOmniorderAllocatedinfoSyncAPIResponseModel is 分单结果同步给星盘 成功返回结果
type TaobaoOmniorderAllocatedinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_allocatedinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误内容
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderAllocatedinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.Message = ""
}

var poolTaobaoOmniorderAllocatedinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderAllocatedinfoSyncAPIResponse)
	},
}

// GetTaobaoOmniorderAllocatedinfoSyncAPIResponse 从 sync.Pool 获取 TaobaoOmniorderAllocatedinfoSyncAPIResponse
func GetTaobaoOmniorderAllocatedinfoSyncAPIResponse() *TaobaoOmniorderAllocatedinfoSyncAPIResponse {
	return poolTaobaoOmniorderAllocatedinfoSyncAPIResponse.Get().(*TaobaoOmniorderAllocatedinfoSyncAPIResponse)
}

// ReleaseTaobaoOmniorderAllocatedinfoSyncAPIResponse 将 TaobaoOmniorderAllocatedinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderAllocatedinfoSyncAPIResponse(v *TaobaoOmniorderAllocatedinfoSyncAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderAllocatedinfoSyncAPIResponse.Put(v)
}
