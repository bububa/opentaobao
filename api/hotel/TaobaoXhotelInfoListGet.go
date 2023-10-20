package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// Taobaoxhotelinfolistget 酒店详细信息查询
// taobao.xhotel.info.list.get
//
// 获取酒店详情信息
func Taobaoxhotelinfolistget(clt *core.SDKClient, req *hotel.TaobaoxhotelinfolistgetAPIRequest, session string) (*hotel.TaobaoxhotelinfolistgetAPIResponse, error) {
	var resp hotel.TaobaoxhotelinfolistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
