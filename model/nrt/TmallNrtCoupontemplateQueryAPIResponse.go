package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCoupontemplateQueryAPIResponse 券模板查询 API返回值
// tmall.nrt.coupontemplate.query
//
// 新零售场景，商家拉取在新零售工作台设置的券数据
type TmallNrtCoupontemplateQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtCoupontemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtCoupontemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtCoupontemplateQueryAPIResponseModel).Reset()
}

// TmallNrtCoupontemplateQueryAPIResponseModel is 券模板查询 成功返回结果
type TmallNrtCoupontemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupontemplate_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// model
	Model *PageData `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtCoupontemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Model = nil
}

var poolTmallNrtCoupontemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtCoupontemplateQueryAPIResponse)
	},
}

// GetTmallNrtCoupontemplateQueryAPIResponse 从 sync.Pool 获取 TmallNrtCoupontemplateQueryAPIResponse
func GetTmallNrtCoupontemplateQueryAPIResponse() *TmallNrtCoupontemplateQueryAPIResponse {
	return poolTmallNrtCoupontemplateQueryAPIResponse.Get().(*TmallNrtCoupontemplateQueryAPIResponse)
}

// ReleaseTmallNrtCoupontemplateQueryAPIResponse 将 TmallNrtCoupontemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtCoupontemplateQueryAPIResponse(v *TmallNrtCoupontemplateQueryAPIResponse) {
	v.Reset()
	poolTmallNrtCoupontemplateQueryAPIResponse.Put(v)
}
