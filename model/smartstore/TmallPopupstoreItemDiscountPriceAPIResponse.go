package smartstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreItemDiscountPriceAPIResponse 商品优惠价格查询 API返回值
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
type TmallPopupstoreItemDiscountPriceAPIResponse struct {
	model.CommonResponse
	TmallPopupstoreItemDiscountPriceAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPopupstoreItemDiscountPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPopupstoreItemDiscountPriceAPIResponseModel).Reset()
}

// TmallPopupstoreItemDiscountPriceAPIResponseModel is 商品优惠价格查询 成功返回结果
type TmallPopupstoreItemDiscountPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_popupstore_item_discount_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	ResultDto *TmallPopupstoreItemDiscountPriceResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}

// Reset 清空结构体
func (m *TmallPopupstoreItemDiscountPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultDto = nil
}

var poolTmallPopupstoreItemDiscountPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreItemDiscountPriceAPIResponse)
	},
}

// GetTmallPopupstoreItemDiscountPriceAPIResponse 从 sync.Pool 获取 TmallPopupstoreItemDiscountPriceAPIResponse
func GetTmallPopupstoreItemDiscountPriceAPIResponse() *TmallPopupstoreItemDiscountPriceAPIResponse {
	return poolTmallPopupstoreItemDiscountPriceAPIResponse.Get().(*TmallPopupstoreItemDiscountPriceAPIResponse)
}

// ReleaseTmallPopupstoreItemDiscountPriceAPIResponse 将 TmallPopupstoreItemDiscountPriceAPIResponse 保存到 sync.Pool
func ReleaseTmallPopupstoreItemDiscountPriceAPIResponse(v *TmallPopupstoreItemDiscountPriceAPIResponse) {
	v.Reset()
	poolTmallPopupstoreItemDiscountPriceAPIResponse.Put(v)
}
