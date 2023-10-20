package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabaalicomtradeadvertiseinfoget 获取订单上的在信息流投放信息
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
func Alibabaalicomtradeadvertiseinfoget(clt *core.SDKClient, req *trade.AlibabaalicomtradeadvertiseinfogetAPIRequest, session string) (*trade.AlibabaalicomtradeadvertiseinfogetAPIResponse, error) {
	var resp trade.AlibabaalicomtradeadvertiseinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
