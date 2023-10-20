package smartstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallpopupstoreitemdiscountpriceAPIResponse 商品优惠价格查询 API返回值
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
type TmallpopupstoreitemdiscountpriceAPIResponse struct {
	model.CommonResponse
	TmallpopupstoreitemdiscountpriceAPIResponseModel
}

// TmallpopupstoreitemdiscountpriceAPIResponseModel is 商品优惠价格查询 成功返回结果
type TmallpopupstoreitemdiscountpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_popupstore_item_discount_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	ResultDto *TmallpopupstoreitemdiscountpriceResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}
