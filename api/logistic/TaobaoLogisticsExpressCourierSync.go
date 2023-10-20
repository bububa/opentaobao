package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressCourierSync 快递公司同步小件员信息
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
func TaobaoLogisticsExpressCourierSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressCourierSyncAPIRequest, session string) (*logistic.TaobaoLogisticsExpressCourierSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressCourierSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
