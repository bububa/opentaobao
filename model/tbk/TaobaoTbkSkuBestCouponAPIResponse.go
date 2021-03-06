package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkSkuBestCouponAPIResponse sku维度最优优惠券信息 API返回值
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
type TaobaoTbkSkuBestCouponAPIResponse struct {
	model.CommonResponse
	TaobaoTbkSkuBestCouponAPIResponseModel
}

// TaobaoTbkSkuBestCouponAPIResponseModel is sku维度最优优惠券信息 成功返回结果
type TaobaoTbkSkuBestCouponAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sku_best_coupon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券详细信息
	CouponInfo *TaobaoTbkSkuBestCouponMapData `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
}
