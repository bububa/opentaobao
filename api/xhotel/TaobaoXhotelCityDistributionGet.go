package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhotelcitydistributionget 酒店城市数据获取接口-分销场景使用
// taobao.xhotel.city.distribution.get
//
// 引流API，对外提供酒店城市数据
func Taobaoxhotelcitydistributionget(clt *core.SDKClient, req *xhotel.TaobaoxhotelcitydistributiongetAPIRequest, session string) (*xhotel.TaobaoxhotelcitydistributiongetAPIResponse, error) {
	var resp xhotel.TaobaoxhotelcitydistributiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
