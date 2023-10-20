package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaordcaligeniusordermsgupdate 订单消息状态回传
// taobao.rdc.aligenius.ordermsg.update
//
// 用于订单消息处理状态回传
func Taobaordcaligeniusordermsgupdate(clt *core.SDKClient, req *trade.TaobaordcaligeniusordermsgupdateAPIRequest, session string) (*trade.TaobaordcaligeniusordermsgupdateAPIResponse, error) {
	var resp trade.TaobaordcaligeniusordermsgupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
