package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorItemsGetAPIResponse 查询商品下载记录 API返回值
// taobao.fenxiao.distributor.items.get
//
// 供应商查询分销商商品下载记录。
type TaobaoFenxiaoDistributorItemsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDistributorItemsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorItemsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDistributorItemsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoDistributorItemsGetAPIResponseModel is 查询商品下载记录 成功返回结果
type TaobaoFenxiaoDistributorItemsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_distributor_items_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下载记录对象
	Records []FenxiaoItemRecord `json:"records,omitempty" xml:"records>fenxiao_item_record,omitempty"`
	// 查询结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorItemsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Records = m.Records[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoDistributorItemsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDistributorItemsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoDistributorItemsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDistributorItemsGetAPIResponse
func GetTaobaoFenxiaoDistributorItemsGetAPIResponse() *TaobaoFenxiaoDistributorItemsGetAPIResponse {
	return poolTaobaoFenxiaoDistributorItemsGetAPIResponse.Get().(*TaobaoFenxiaoDistributorItemsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoDistributorItemsGetAPIResponse 将 TaobaoFenxiaoDistributorItemsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDistributorItemsGetAPIResponse(v *TaobaoFenxiaoDistributorItemsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDistributorItemsGetAPIResponse.Put(v)
}
