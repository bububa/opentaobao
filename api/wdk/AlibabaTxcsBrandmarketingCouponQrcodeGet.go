package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTxcsBrandmarketingCouponQrcodeGet 品牌营销导购员券页面二维码获取
// alibaba.txcs.brandmarketing.coupon.qrcode.get
//
// 构建券页码二维码url
func AlibabaTxcsBrandmarketingCouponQrcodeGet(clt *core.SDKClient, req *wdk.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest, resp *wdk.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
