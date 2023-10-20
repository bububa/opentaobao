package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpresscollectsync 服饰逆向揽收信息同步
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
func Taobaologisticsexpresscollectsync(clt *core.SDKClient, req *logistic.TaobaologisticsexpresscollectsyncAPIRequest, session string) (*logistic.TaobaologisticsexpresscollectsyncAPIResponse, error) {
	var resp logistic.TaobaologisticsexpresscollectsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
