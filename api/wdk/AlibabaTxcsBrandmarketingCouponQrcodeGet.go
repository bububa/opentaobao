package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatxcsbrandmarketingcouponqrcodeget 品牌营销导购员券页面二维码获取
// alibaba.txcs.brandmarketing.coupon.qrcode.get
//
// 构建券页码二维码url
func Alibabatxcsbrandmarketingcouponqrcodeget(clt *core.SDKClient, req *wdk.AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest, session string) (*wdk.AlibabatxcsbrandmarketingcouponqrcodegetAPIResponse, error) {
	var resp wdk.AlibabatxcsbrandmarketingcouponqrcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
