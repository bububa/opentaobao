package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpaypromotioninfoget tv支付查询消费抽奖配置
// taobao.tvpay.promotion.info.get
//
// 查询消费抽奖配置
func Taobaotvpaypromotioninfoget(clt *core.SDKClient, req *tvpay.TaobaotvpaypromotioninfogetAPIRequest, session string) (*tvpay.TaobaotvpaypromotioninfogetAPIResponse, error) {
	var resp tvpay.TaobaotvpaypromotioninfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
