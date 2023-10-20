package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahtcouponfuluphonechargecallback 话费充值回调
// alibaba.htcoupon.fulu.phonecharge.callback
//
// 话费充值为异步操作，此接口用于充值成功后，供应商回调。
func Alibabahtcouponfuluphonechargecallback(clt *core.SDKClient, req *happytrip.AlibabahtcouponfuluphonechargecallbackAPIRequest, session string) (*happytrip.AlibabahtcouponfuluphonechargecallbackAPIResponse, error) {
	var resp happytrip.AlibabahtcouponfuluphonechargecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
