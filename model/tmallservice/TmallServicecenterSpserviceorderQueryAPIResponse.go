package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderQueryAPIResponse 服务单列表查询 API返回值
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
type TmallServicecenterSpserviceorderQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterSpserviceorderQueryAPIResponseModel).Reset()
}

// TmallServicecenterSpserviceorderQueryAPIResponseModel is 服务单列表查询 成功返回结果
type TmallServicecenterSpserviceorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterSpserviceorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterSpserviceorderQueryAPIResponse)
	},
}

// GetTmallServicecenterSpserviceorderQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterSpserviceorderQueryAPIResponse
func GetTmallServicecenterSpserviceorderQueryAPIResponse() *TmallServicecenterSpserviceorderQueryAPIResponse {
	return poolTmallServicecenterSpserviceorderQueryAPIResponse.Get().(*TmallServicecenterSpserviceorderQueryAPIResponse)
}

// ReleaseTmallServicecenterSpserviceorderQueryAPIResponse 将 TmallServicecenterSpserviceorderQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterSpserviceorderQueryAPIResponse(v *TmallServicecenterSpserviceorderQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterSpserviceorderQueryAPIResponse.Put(v)
}
