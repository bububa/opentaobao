package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrZqsPlanQueryAPIResponse 周期送配送明细查询 API返回值
// tmall.nr.zqs.plan.query
//
// 周期送配送明细查询
type TmallNrZqsPlanQueryAPIResponse struct {
	model.CommonResponse
	TmallNrZqsPlanQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrZqsPlanQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrZqsPlanQueryAPIResponseModel).Reset()
}

// TmallNrZqsPlanQueryAPIResponseModel is 周期送配送明细查询 成功返回结果
type TmallNrZqsPlanQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_zqs_plan_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrZqsPlanQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrZqsPlanQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrZqsPlanQueryAPIResponse)
	},
}

// GetTmallNrZqsPlanQueryAPIResponse 从 sync.Pool 获取 TmallNrZqsPlanQueryAPIResponse
func GetTmallNrZqsPlanQueryAPIResponse() *TmallNrZqsPlanQueryAPIResponse {
	return poolTmallNrZqsPlanQueryAPIResponse.Get().(*TmallNrZqsPlanQueryAPIResponse)
}

// ReleaseTmallNrZqsPlanQueryAPIResponse 将 TmallNrZqsPlanQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrZqsPlanQueryAPIResponse(v *TmallNrZqsPlanQueryAPIResponse) {
	v.Reset()
	poolTmallNrZqsPlanQueryAPIResponse.Put(v)
}
