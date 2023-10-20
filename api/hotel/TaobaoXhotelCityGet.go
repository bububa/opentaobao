package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelCityGet 酒店城市数据获取接口
// taobao.xhotel.city.get
//
// 引流API，对外提供酒店城市数据
func TaobaoXhotelCityGet(clt *core.SDKClient, req *hotel.TaobaoXhotelCityGetAPIRequest, resp *hotel.TaobaoXhotelCityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
