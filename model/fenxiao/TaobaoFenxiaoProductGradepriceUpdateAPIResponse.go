package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductGradepriceUpdateAPIResponse 根据sku设置折扣价 API返回值
// taobao.fenxiao.product.gradeprice.update
//
// 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
type TaobaoFenxiaoProductGradepriceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductGradepriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductGradepriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductGradepriceUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductGradepriceUpdateAPIResponseModel is 根据sku设置折扣价 成功返回结果
type TaobaoFenxiaoProductGradepriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_gradeprice_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回操作结果：成功或失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductGradepriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductGradepriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductGradepriceUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoProductGradepriceUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductGradepriceUpdateAPIResponse
func GetTaobaoFenxiaoProductGradepriceUpdateAPIResponse() *TaobaoFenxiaoProductGradepriceUpdateAPIResponse {
	return poolTaobaoFenxiaoProductGradepriceUpdateAPIResponse.Get().(*TaobaoFenxiaoProductGradepriceUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoProductGradepriceUpdateAPIResponse 将 TaobaoFenxiaoProductGradepriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductGradepriceUpdateAPIResponse(v *TaobaoFenxiaoProductGradepriceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductGradepriceUpdateAPIResponse.Put(v)
}
