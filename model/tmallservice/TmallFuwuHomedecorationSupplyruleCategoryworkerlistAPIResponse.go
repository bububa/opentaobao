package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse 基于规则查品牌品类工人接口 API返回值
// tmall.fuwu.homedecoration.supplyrule.categoryworkerlist
//
// 基于规则查品牌品类工人接口
type TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse struct {
	model.CommonResponse
	TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponseModel).Reset()
}

// TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponseModel is 基于规则查品牌品类工人接口 成功返回结果
type TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_homedecoration_supplyrule_categoryworkerlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse)
	},
}

// GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse 从 sync.Pool 获取 TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse
func GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse() *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse {
	return poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse.Get().(*TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse)
}

// ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse 将 TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse 保存到 sync.Pool
func ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse(v *TmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse) {
	v.Reset()
	poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistAPIResponse.Put(v)
}
