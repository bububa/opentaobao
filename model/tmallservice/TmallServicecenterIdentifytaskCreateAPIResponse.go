package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterIdentifytaskCreateAPIResponse 服务商创建核销单 API返回值
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
type TmallServicecenterIdentifytaskCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterIdentifytaskCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterIdentifytaskCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterIdentifytaskCreateAPIResponseModel).Reset()
}

// TmallServicecenterIdentifytaskCreateAPIResponseModel is 服务商创建核销单 成功返回结果
type TmallServicecenterIdentifytaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_identifytask_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// -
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterIdentifytaskCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterIdentifytaskCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterIdentifytaskCreateAPIResponse)
	},
}

// GetTmallServicecenterIdentifytaskCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterIdentifytaskCreateAPIResponse
func GetTmallServicecenterIdentifytaskCreateAPIResponse() *TmallServicecenterIdentifytaskCreateAPIResponse {
	return poolTmallServicecenterIdentifytaskCreateAPIResponse.Get().(*TmallServicecenterIdentifytaskCreateAPIResponse)
}

// ReleaseTmallServicecenterIdentifytaskCreateAPIResponse 将 TmallServicecenterIdentifytaskCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterIdentifytaskCreateAPIResponse(v *TmallServicecenterIdentifytaskCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterIdentifytaskCreateAPIResponse.Put(v)
}
