package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationSupplyruleListAPIResponse 查询供给规则接口 API返回值
// tmall.fuwu.homedecoration.supplyrule.list
//
// 查询供给规则接口
type TmallFuwuHomedecorationSupplyruleListAPIResponse struct {
	model.CommonResponse
	TmallFuwuHomedecorationSupplyruleListAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFuwuHomedecorationSupplyruleListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFuwuHomedecorationSupplyruleListAPIResponseModel).Reset()
}

// TmallFuwuHomedecorationSupplyruleListAPIResponseModel is 查询供给规则接口 成功返回结果
type TmallFuwuHomedecorationSupplyruleListAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_homedecoration_supplyrule_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallFuwuHomedecorationSupplyruleListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFuwuHomedecorationSupplyruleListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallFuwuHomedecorationSupplyruleListAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFuwuHomedecorationSupplyruleListAPIResponse)
	},
}

// GetTmallFuwuHomedecorationSupplyruleListAPIResponse 从 sync.Pool 获取 TmallFuwuHomedecorationSupplyruleListAPIResponse
func GetTmallFuwuHomedecorationSupplyruleListAPIResponse() *TmallFuwuHomedecorationSupplyruleListAPIResponse {
	return poolTmallFuwuHomedecorationSupplyruleListAPIResponse.Get().(*TmallFuwuHomedecorationSupplyruleListAPIResponse)
}

// ReleaseTmallFuwuHomedecorationSupplyruleListAPIResponse 将 TmallFuwuHomedecorationSupplyruleListAPIResponse 保存到 sync.Pool
func ReleaseTmallFuwuHomedecorationSupplyruleListAPIResponse(v *TmallFuwuHomedecorationSupplyruleListAPIResponse) {
	v.Reset()
	poolTmallFuwuHomedecorationSupplyruleListAPIResponse.Put(v)
}
