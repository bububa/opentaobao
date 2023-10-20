package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabatradealiancecreate 推客平台订单回流
// alibaba.trade.aliance.create
//
// 推客平台订单回流
func Alibabatradealiancecreate(clt *core.SDKClient, req *trade.AlibabatradealiancecreateAPIRequest, session string) (*trade.AlibabatradealiancecreateAPIResponse, error) {
	var resp trade.AlibabatradealiancecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
