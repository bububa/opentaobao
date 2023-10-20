package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtorderHotelSyncConfigAPIResponse 同步配置信息 API返回值
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
type AlibabaHtorderHotelSyncConfigAPIResponse struct {
	model.CommonResponse
	AlibabaHtorderHotelSyncConfigAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHtorderHotelSyncConfigAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHtorderHotelSyncConfigAPIResponseModel).Reset()
}

// AlibabaHtorderHotelSyncConfigAPIResponseModel is 同步配置信息 成功返回结果
type AlibabaHtorderHotelSyncConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_htorder_hotel_sync_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrNo string `json:"err_no,omitempty" xml:"err_no,omitempty"`
	// 栈信息
	StackTrace string `json:"stack_trace,omitempty" xml:"stack_trace,omitempty"`
	// 错误信息
	ErrInfo string `json:"err_info,omitempty" xml:"err_info,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 成功OR失败
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHtorderHotelSyncConfigAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrNo = ""
	m.StackTrace = ""
	m.ErrInfo = ""
	m.Succ = false
	m.Content = false
}

var poolAlibabaHtorderHotelSyncConfigAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHtorderHotelSyncConfigAPIResponse)
	},
}

// GetAlibabaHtorderHotelSyncConfigAPIResponse 从 sync.Pool 获取 AlibabaHtorderHotelSyncConfigAPIResponse
func GetAlibabaHtorderHotelSyncConfigAPIResponse() *AlibabaHtorderHotelSyncConfigAPIResponse {
	return poolAlibabaHtorderHotelSyncConfigAPIResponse.Get().(*AlibabaHtorderHotelSyncConfigAPIResponse)
}

// ReleaseAlibabaHtorderHotelSyncConfigAPIResponse 将 AlibabaHtorderHotelSyncConfigAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHtorderHotelSyncConfigAPIResponse(v *AlibabaHtorderHotelSyncConfigAPIResponse) {
	v.Reset()
	poolAlibabaHtorderHotelSyncConfigAPIResponse.Put(v)
}
