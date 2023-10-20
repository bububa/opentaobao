package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageSearchAPIResponse 根据时间段查询服务商的服务预警消息列表(15分钟内) API返回值
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
type TmallServicecenterServicemonitormessageSearchAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicemonitormessageSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicemonitormessageSearchAPIResponseModel).Reset()
}

// TmallServicecenterServicemonitormessageSearchAPIResponseModel is 根据时间段查询服务商的服务预警消息列表(15分钟内) 成功返回结果
type TmallServicecenterServicemonitormessageSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicemonitormessage_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicemonitormessageSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicemonitormessageSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicemonitormessageSearchAPIResponse)
	},
}

// GetTmallServicecenterServicemonitormessageSearchAPIResponse 从 sync.Pool 获取 TmallServicecenterServicemonitormessageSearchAPIResponse
func GetTmallServicecenterServicemonitormessageSearchAPIResponse() *TmallServicecenterServicemonitormessageSearchAPIResponse {
	return poolTmallServicecenterServicemonitormessageSearchAPIResponse.Get().(*TmallServicecenterServicemonitormessageSearchAPIResponse)
}

// ReleaseTmallServicecenterServicemonitormessageSearchAPIResponse 将 TmallServicecenterServicemonitormessageSearchAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageSearchAPIResponse(v *TmallServicecenterServicemonitormessageSearchAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageSearchAPIResponse.Put(v)
}
