package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingcouponsendma 发放匿名码
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
func Alibabahmmarketingcouponsendma(clt *core.SDKClient, req *wdk.AlibabahmmarketingcouponsendmaAPIRequest, session string) (*wdk.AlibabahmmarketingcouponsendmaAPIResponse, error) {
	var resp wdk.AlibabahmmarketingcouponsendmaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
