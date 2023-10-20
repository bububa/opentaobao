package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrGameDataSyncAPIResponse 公益互动 外部游戏数据同步 API返回值
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
type AlibabaCsrGameDataSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCsrGameDataSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrGameDataSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrGameDataSyncAPIResponseModel).Reset()
}

// AlibabaCsrGameDataSyncAPIResponseModel is 公益互动 外部游戏数据同步 成功返回结果
type AlibabaCsrGameDataSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_game_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误内容
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrGameDataSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = 0
	m.ResultSuccess = false
}

var poolAlibabaCsrGameDataSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrGameDataSyncAPIResponse)
	},
}

// GetAlibabaCsrGameDataSyncAPIResponse 从 sync.Pool 获取 AlibabaCsrGameDataSyncAPIResponse
func GetAlibabaCsrGameDataSyncAPIResponse() *AlibabaCsrGameDataSyncAPIResponse {
	return poolAlibabaCsrGameDataSyncAPIResponse.Get().(*AlibabaCsrGameDataSyncAPIResponse)
}

// ReleaseAlibabaCsrGameDataSyncAPIResponse 将 AlibabaCsrGameDataSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrGameDataSyncAPIResponse(v *AlibabaCsrGameDataSyncAPIResponse) {
	v.Reset()
	poolAlibabaCsrGameDataSyncAPIResponse.Put(v)
}
