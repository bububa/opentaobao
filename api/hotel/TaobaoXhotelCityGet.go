package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// Taobaoxhotelcityget 酒店城市数据获取接口
// taobao.xhotel.city.get
//
// 引流API，对外提供酒店城市数据
func Taobaoxhotelcityget(clt *core.SDKClient, req *hotel.TaobaoxhotelcitygetAPIRequest, session string) (*hotel.TaobaoxhotelcitygetAPIResponse, error) {
	var resp hotel.TaobaoxhotelcitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
