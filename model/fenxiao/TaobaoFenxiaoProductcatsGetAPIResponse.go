package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatsGetAPIResponse 查询产品线列表 API返回值
// taobao.fenxiao.productcats.get
//
// 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductcatsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductcatsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductcatsGetAPIResponseModel is 查询产品线列表 成功返回结果
type TaobaoFenxiaoProductcatsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_productcats_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品线列表。返回 ProductCat 包含的字段信息。
	Productcats []ProductCat `json:"productcats,omitempty" xml:"productcats>product_cat,omitempty"`
	// 查询结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Productcats = m.Productcats[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoProductcatsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductcatsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoProductcatsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductcatsGetAPIResponse
func GetTaobaoFenxiaoProductcatsGetAPIResponse() *TaobaoFenxiaoProductcatsGetAPIResponse {
	return poolTaobaoFenxiaoProductcatsGetAPIResponse.Get().(*TaobaoFenxiaoProductcatsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoProductcatsGetAPIResponse 将 TaobaoFenxiaoProductcatsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductcatsGetAPIResponse(v *TaobaoFenxiaoProductcatsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductcatsGetAPIResponse.Put(v)
}
