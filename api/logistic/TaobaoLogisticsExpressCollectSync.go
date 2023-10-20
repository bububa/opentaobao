package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressCollectSync 服饰逆向揽收信息同步
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
func TaobaoLogisticsExpressCollectSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressCollectSyncAPIRequest, session string) (*logistic.TaobaoLogisticsExpressCollectSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressCollectSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
