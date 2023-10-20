package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpressdeliverycutnotify TMS配拦截结果回告
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
func Taobaologisticsexpressdeliverycutnotify(clt *core.SDKClient, req *logistic.TaobaologisticsexpressdeliverycutnotifyAPIRequest, session string) (*logistic.TaobaologisticsexpressdeliverycutnotifyAPIResponse, error) {
	var resp logistic.TaobaologisticsexpressdeliverycutnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
