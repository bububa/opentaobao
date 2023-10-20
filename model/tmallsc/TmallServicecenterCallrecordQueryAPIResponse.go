package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterCallrecordQueryAPIResponse 天猫服务平台服务商查询通话记录接口 API返回值
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
type TmallServicecenterCallrecordQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterCallrecordQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterCallrecordQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterCallrecordQueryAPIResponseModel).Reset()
}

// TmallServicecenterCallrecordQueryAPIResponseModel is 天猫服务平台服务商查询通话记录接口 成功返回结果
type TmallServicecenterCallrecordQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_callrecord_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallServicecenterCallrecordQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterCallrecordQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterCallrecordQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterCallrecordQueryAPIResponse)
	},
}

// GetTmallServicecenterCallrecordQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterCallrecordQueryAPIResponse
func GetTmallServicecenterCallrecordQueryAPIResponse() *TmallServicecenterCallrecordQueryAPIResponse {
	return poolTmallServicecenterCallrecordQueryAPIResponse.Get().(*TmallServicecenterCallrecordQueryAPIResponse)
}

// ReleaseTmallServicecenterCallrecordQueryAPIResponse 将 TmallServicecenterCallrecordQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterCallrecordQueryAPIResponse(v *TmallServicecenterCallrecordQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterCallrecordQueryAPIResponse.Put(v)
}
