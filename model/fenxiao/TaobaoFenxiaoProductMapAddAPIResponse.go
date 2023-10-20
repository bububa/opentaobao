package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductMapAddAPIResponse 创建分销和后端商品映射关系 API返回值
// taobao.fenxiao.product.map.add
//
// 创建分销和供应链商品映射关系。
type TaobaoFenxiaoProductMapAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductMapAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductMapAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductMapAddAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductMapAddAPIResponseModel is 创建分销和后端商品映射关系 成功返回结果
type TaobaoFenxiaoProductMapAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_map_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductMapAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductMapAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductMapAddAPIResponse)
	},
}

// GetTaobaoFenxiaoProductMapAddAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductMapAddAPIResponse
func GetTaobaoFenxiaoProductMapAddAPIResponse() *TaobaoFenxiaoProductMapAddAPIResponse {
	return poolTaobaoFenxiaoProductMapAddAPIResponse.Get().(*TaobaoFenxiaoProductMapAddAPIResponse)
}

// ReleaseTaobaoFenxiaoProductMapAddAPIResponse 将 TaobaoFenxiaoProductMapAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductMapAddAPIResponse(v *TaobaoFenxiaoProductMapAddAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductMapAddAPIResponse.Put(v)
}
