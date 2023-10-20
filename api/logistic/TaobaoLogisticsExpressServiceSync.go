package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressServiceSync 服务信息回告接口
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
func TaobaoLogisticsExpressServiceSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressServiceSyncAPIRequest, session string) (*logistic.TaobaoLogisticsExpressServiceSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsExpressServiceSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
