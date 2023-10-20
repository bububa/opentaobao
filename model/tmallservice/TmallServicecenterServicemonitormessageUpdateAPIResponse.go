package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageUpdateAPIResponse 服务商更新预警消息状态 API返回值
// tmall.servicecenter.servicemonitormessage.update
//
// 服务商收到预警后，需要进行回复已读状态，并可填写备注
type TmallServicecenterServicemonitormessageUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicemonitormessageUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicemonitormessageUpdateAPIResponseModel).Reset()
}

// TmallServicecenterServicemonitormessageUpdateAPIResponseModel is 服务商更新预警消息状态 成功返回结果
type TmallServicecenterServicemonitormessageUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicemonitormessageUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicemonitormessageUpdateAPIResponse)
	},
}

// GetTmallServicecenterServicemonitormessageUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterServicemonitormessageUpdateAPIResponse
func GetTmallServicecenterServicemonitormessageUpdateAPIResponse() *TmallServicecenterServicemonitormessageUpdateAPIResponse {
	return poolTmallServicecenterServicemonitormessageUpdateAPIResponse.Get().(*TmallServicecenterServicemonitormessageUpdateAPIResponse)
}

// ReleaseTmallServicecenterServicemonitormessageUpdateAPIResponse 将 TmallServicecenterServicemonitormessageUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageUpdateAPIResponse(v *TmallServicecenterServicemonitormessageUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageUpdateAPIResponse.Put(v)
}
