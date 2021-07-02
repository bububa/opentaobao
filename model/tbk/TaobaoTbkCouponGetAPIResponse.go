package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkCouponGetAPIResponse 淘宝客-公用-阿里妈妈推广券详情查询 API返回值
// taobao.tbk.coupon.get
//
// 传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
type TaobaoTbkCouponGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkCouponGetAPIResponseModel
}

// TaobaoTbkCouponGetAPIResponseModel is 淘宝客-公用-阿里妈妈推广券详情查询 成功返回结果
type TaobaoTbkCouponGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_coupon_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkCouponGetMapData `json:"data,omitempty" xml:"data,omitempty"`
}
