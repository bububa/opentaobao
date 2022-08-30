package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinAlipayDksign 支付宝代扣签约
// taobao.alitrip.travel.axin.alipay.dksign
//
// 提供支付宝代扣签约服务
func TaobaoAlitripTravelAxinAlipayDksign(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinAlipayDksignAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinAlipayDksignAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinAlipayDksignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
