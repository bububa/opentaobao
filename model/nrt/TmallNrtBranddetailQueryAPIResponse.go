package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtBranddetailQueryAPIResponse 品牌详情查询 API返回值
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
type TmallNrtBranddetailQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtBranddetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtBranddetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtBranddetailQueryAPIResponseModel).Reset()
}

// TmallNrtBranddetailQueryAPIResponseModel is 品牌详情查询 成功返回结果
type TmallNrtBranddetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_branddetail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtBranddetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTmallNrtBranddetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtBranddetailQueryAPIResponse)
	},
}

// GetTmallNrtBranddetailQueryAPIResponse 从 sync.Pool 获取 TmallNrtBranddetailQueryAPIResponse
func GetTmallNrtBranddetailQueryAPIResponse() *TmallNrtBranddetailQueryAPIResponse {
	return poolTmallNrtBranddetailQueryAPIResponse.Get().(*TmallNrtBranddetailQueryAPIResponse)
}

// ReleaseTmallNrtBranddetailQueryAPIResponse 将 TmallNrtBranddetailQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtBranddetailQueryAPIResponse(v *TmallNrtBranddetailQueryAPIResponse) {
	v.Reset()
	poolTmallNrtBranddetailQueryAPIResponse.Put(v)
}
