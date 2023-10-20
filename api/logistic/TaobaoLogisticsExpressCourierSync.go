package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpresscouriersync 快递公司同步小件员信息
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
func Taobaologisticsexpresscouriersync(clt *core.SDKClient, req *logistic.TaobaologisticsexpresscouriersyncAPIRequest, session string) (*logistic.TaobaologisticsexpresscouriersyncAPIResponse, error) {
	var resp logistic.TaobaologisticsexpresscouriersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
