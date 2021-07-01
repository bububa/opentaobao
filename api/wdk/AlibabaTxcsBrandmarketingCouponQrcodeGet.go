package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
品牌营销导购员券页面二维码获取 
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url
*/
func AlibabaTxcsBrandmarketingCouponQrcodeGet(clt *core.SDKClient, req *wdk.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest, session string) (*wdk.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse, error) {
    var resp wdk.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
