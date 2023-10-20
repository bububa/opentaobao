package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaoopensellerbizlogistictimerule 商家自定义发货时效
// taobao.open.seller.biz.logistic.time.rule
//
// 服务商回传商家自定义发货时效
func Taobaoopensellerbizlogistictimerule(clt *core.SDKClient, req *logistic.TaobaoopensellerbizlogistictimeruleAPIRequest, session string) (*logistic.TaobaoopensellerbizlogistictimeruleAPIResponse, error) {
	var resp logistic.TaobaoopensellerbizlogistictimeruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
