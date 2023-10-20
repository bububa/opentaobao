package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaotradewtverticalget 网厅垂直信息查询接口
// taobao.trade.wtvertical.get
//
// 网厅订单垂直信息的查询
func Taobaotradewtverticalget(clt *core.SDKClient, req *trade.TaobaotradewtverticalgetAPIRequest, session string) (*trade.TaobaotradewtverticalgetAPIResponse, error) {
	var resp trade.TaobaotradewtverticalgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
