package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressCourierSync 快递公司同步小件员信息
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
func TaobaoLogisticsExpressCourierSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressCourierSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressCourierSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
