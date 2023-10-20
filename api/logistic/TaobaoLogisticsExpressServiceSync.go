package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpressservicesync 服务信息回告接口
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
func Taobaologisticsexpressservicesync(clt *core.SDKClient, req *logistic.TaobaologisticsexpressservicesyncAPIRequest, session string) (*logistic.TaobaologisticsexpressservicesyncAPIResponse, error) {
	var resp logistic.TaobaologisticsexpressservicesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
