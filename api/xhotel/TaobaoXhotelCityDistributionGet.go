package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelCityDistributionGet 酒店城市数据获取接口-分销场景使用
// taobao.xhotel.city.distribution.get
//
// 引流API，对外提供酒店城市数据
func TaobaoXhotelCityDistributionGet(clt *core.SDKClient, req *xhotel.TaobaoXhotelCityDistributionGetAPIRequest, resp *xhotel.TaobaoXhotelCityDistributionGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
