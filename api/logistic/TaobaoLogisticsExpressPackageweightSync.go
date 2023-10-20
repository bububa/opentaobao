package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpresspackageweightsync TMS包裹重量回传
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
func Taobaologisticsexpresspackageweightsync(clt *core.SDKClient, req *logistic.TaobaologisticsexpresspackageweightsyncAPIRequest, session string) (*logistic.TaobaologisticsexpresspackageweightsyncAPIResponse, error) {
	var resp logistic.TaobaologisticsexpresspackageweightsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
