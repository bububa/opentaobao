package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaowtttradeserviceget 获取网厅号卡垂直标信息
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
func Taobaowtttradeserviceget(clt *core.SDKClient, req *trade.TaobaowtttradeservicegetAPIRequest, session string) (*trade.TaobaowtttradeservicegetAPIResponse, error) {
	var resp trade.TaobaowtttradeservicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
