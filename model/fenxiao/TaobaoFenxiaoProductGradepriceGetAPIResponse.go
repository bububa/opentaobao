package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductGradepriceGetAPIResponse 等级折扣查询 API返回值
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
type TaobaoFenxiaoProductGradepriceGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductGradepriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductGradepriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductGradepriceGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductGradepriceGetAPIResponseModel is 等级折扣查询 成功返回结果
type TaobaoFenxiaoProductGradepriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_gradeprice_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 等级折扣列表
	GradeDiscounts []GradeDiscount `json:"grade_discounts,omitempty" xml:"grade_discounts>grade_discount,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductGradepriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GradeDiscounts = m.GradeDiscounts[:0]
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductGradepriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductGradepriceGetAPIResponse)
	},
}

// GetTaobaoFenxiaoProductGradepriceGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductGradepriceGetAPIResponse
func GetTaobaoFenxiaoProductGradepriceGetAPIResponse() *TaobaoFenxiaoProductGradepriceGetAPIResponse {
	return poolTaobaoFenxiaoProductGradepriceGetAPIResponse.Get().(*TaobaoFenxiaoProductGradepriceGetAPIResponse)
}

// ReleaseTaobaoFenxiaoProductGradepriceGetAPIResponse 将 TaobaoFenxiaoProductGradepriceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductGradepriceGetAPIResponse(v *TaobaoFenxiaoProductGradepriceGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductGradepriceGetAPIResponse.Put(v)
}
