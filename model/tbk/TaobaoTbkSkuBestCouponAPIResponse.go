package tbk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTbkSkuBestCouponAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkSkuBestCouponAPIResponseModel).Reset()
}

// TaobaoTbkSkuBestCouponAPIResponseModel is sku维度最优优惠券信息 成功返回结果
type TaobaoTbkSkuBestCouponAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sku_best_coupon_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券详细信息
	CouponInfo *TaobaoTbkSkuBestCouponMapData `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkSkuBestCouponAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CouponInfo = nil
}

var poolTaobaoTbkSkuBestCouponAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkSkuBestCouponAPIResponse)
	},
}

// GetTaobaoTbkSkuBestCouponAPIResponse 从 sync.Pool 获取 TaobaoTbkSkuBestCouponAPIResponse
func GetTaobaoTbkSkuBestCouponAPIResponse() *TaobaoTbkSkuBestCouponAPIResponse {
	return poolTaobaoTbkSkuBestCouponAPIResponse.Get().(*TaobaoTbkSkuBestCouponAPIResponse)
}

// ReleaseTaobaoTbkSkuBestCouponAPIResponse 将 TaobaoTbkSkuBestCouponAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkSkuBestCouponAPIResponse(v *TaobaoTbkSkuBestCouponAPIResponse) {
	v.Reset()
	poolTaobaoTbkSkuBestCouponAPIResponse.Put(v)
}
