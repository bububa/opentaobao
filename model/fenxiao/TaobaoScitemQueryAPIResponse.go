package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemQueryAPIResponse 查询后端商品 API返回值
// taobao.scitem.query
//
// 查询后端商品
type TaobaoScitemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoScitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemQueryAPIResponseModel).Reset()
}

// TaobaoScitemQueryAPIResponseModel is 查询后端商品 成功返回结果
type TaobaoScitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// List&lt;ScItemDO&gt;
	ScItemList []ScItem `json:"sc_item_list,omitempty" xml:"sc_item_list>sc_item,omitempty"`
	// 商品条数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 分页
	QueryPagination *QueryPagination `json:"query_pagination,omitempty" xml:"query_pagination,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScItemList = m.ScItemList[:0]
	m.TotalPage = 0
	m.QueryPagination = nil
}

var poolTaobaoScitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemQueryAPIResponse)
	},
}

// GetTaobaoScitemQueryAPIResponse 从 sync.Pool 获取 TaobaoScitemQueryAPIResponse
func GetTaobaoScitemQueryAPIResponse() *TaobaoScitemQueryAPIResponse {
	return poolTaobaoScitemQueryAPIResponse.Get().(*TaobaoScitemQueryAPIResponse)
}

// ReleaseTaobaoScitemQueryAPIResponse 将 TaobaoScitemQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemQueryAPIResponse(v *TaobaoScitemQueryAPIResponse) {
	v.Reset()
	poolTaobaoScitemQueryAPIResponse.Put(v)
}
