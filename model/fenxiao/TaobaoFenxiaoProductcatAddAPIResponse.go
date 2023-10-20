package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatAddAPIResponse 新增产品线 API返回值
// taobao.fenxiao.productcat.add
//
// 新增产品线
type TaobaoFenxiaoProductcatAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductcatAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductcatAddAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductcatAddAPIResponseModel is 新增产品线 成功返回结果
type TaobaoFenxiaoProductcatAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_productcat_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品线ID
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductLineId = 0
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductcatAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductcatAddAPIResponse)
	},
}

// GetTaobaoFenxiaoProductcatAddAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductcatAddAPIResponse
func GetTaobaoFenxiaoProductcatAddAPIResponse() *TaobaoFenxiaoProductcatAddAPIResponse {
	return poolTaobaoFenxiaoProductcatAddAPIResponse.Get().(*TaobaoFenxiaoProductcatAddAPIResponse)
}

// ReleaseTaobaoFenxiaoProductcatAddAPIResponse 将 TaobaoFenxiaoProductcatAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductcatAddAPIResponse(v *TaobaoFenxiaoProductcatAddAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductcatAddAPIResponse.Put(v)
}
