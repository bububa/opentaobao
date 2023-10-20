package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaoopensellerbizlogisticsellerbind 店铺授权发货注册（催发货）
// taobao.open.seller.biz.logistic.seller.bind
//
// 店铺授权发货注册（催发货）
func Taobaoopensellerbizlogisticsellerbind(clt *core.SDKClient, req *logistic.TaobaoopensellerbizlogisticsellerbindAPIRequest, session string) (*logistic.TaobaoopensellerbizlogisticsellerbindAPIResponse, error) {
	var resp logistic.TaobaoopensellerbizlogisticsellerbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
