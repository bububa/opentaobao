package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaReportCityGet 获取城市维度报表
// taobao.simba.report.city.get
//
// 获取城市维度报表
func TaobaoSimbaReportCityGet(clt *core.SDKClient, req *simba.TaobaoSimbaReportCityGetAPIRequest, resp *simba.TaobaoSimbaReportCityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
