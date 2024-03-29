package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrGameDataSyncCheckAPIResponse 公益互动 外部游戏数据同步-校验 API返回值
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
type AlibabaCsrGameDataSyncCheckAPIResponse struct {
	model.CommonResponse
	AlibabaCsrGameDataSyncCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrGameDataSyncCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrGameDataSyncCheckAPIResponseModel).Reset()
}

// AlibabaCsrGameDataSyncCheckAPIResponseModel is 公益互动 外部游戏数据同步-校验 成功返回结果
type AlibabaCsrGameDataSyncCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_game_data_sync_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrGameDataSyncCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = 0
	m.ResultSuccess = false
}

var poolAlibabaCsrGameDataSyncCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrGameDataSyncCheckAPIResponse)
	},
}

// GetAlibabaCsrGameDataSyncCheckAPIResponse 从 sync.Pool 获取 AlibabaCsrGameDataSyncCheckAPIResponse
func GetAlibabaCsrGameDataSyncCheckAPIResponse() *AlibabaCsrGameDataSyncCheckAPIResponse {
	return poolAlibabaCsrGameDataSyncCheckAPIResponse.Get().(*AlibabaCsrGameDataSyncCheckAPIResponse)
}

// ReleaseAlibabaCsrGameDataSyncCheckAPIResponse 将 AlibabaCsrGameDataSyncCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrGameDataSyncCheckAPIResponse(v *AlibabaCsrGameDataSyncCheckAPIResponse) {
	v.Reset()
	poolAlibabaCsrGameDataSyncCheckAPIResponse.Put(v)
}
