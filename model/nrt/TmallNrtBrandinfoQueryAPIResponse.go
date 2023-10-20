package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtBrandinfoQueryAPIResponse 品牌数据查询 API返回值
// tmall.nrt.brandinfo.query
//
// 商家获取自己旗舰店授权的品牌id列表
type TmallNrtBrandinfoQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtBrandinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtBrandinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtBrandinfoQueryAPIResponseModel).Reset()
}

// TmallNrtBrandinfoQueryAPIResponseModel is 品牌数据查询 成功返回结果
type TmallNrtBrandinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_brandinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 品牌id
	Datas []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtBrandinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTmallNrtBrandinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtBrandinfoQueryAPIResponse)
	},
}

// GetTmallNrtBrandinfoQueryAPIResponse 从 sync.Pool 获取 TmallNrtBrandinfoQueryAPIResponse
func GetTmallNrtBrandinfoQueryAPIResponse() *TmallNrtBrandinfoQueryAPIResponse {
	return poolTmallNrtBrandinfoQueryAPIResponse.Get().(*TmallNrtBrandinfoQueryAPIResponse)
}

// ReleaseTmallNrtBrandinfoQueryAPIResponse 将 TmallNrtBrandinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtBrandinfoQueryAPIResponse(v *TmallNrtBrandinfoQueryAPIResponse) {
	v.Reset()
	poolTmallNrtBrandinfoQueryAPIResponse.Put(v)
}
