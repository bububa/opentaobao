package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripmonthbillurlget 月账单数据查询
// alitrip.btrip.monthbill.url.get
//
// 月账单数据查询
func Alitripbtripmonthbillurlget(clt *core.SDKClient, req *btrip.AlitripbtripmonthbillurlgetAPIRequest, session string) (*btrip.AlitripbtripmonthbillurlgetAPIResponse, error) {
	var resp btrip.AlitripbtripmonthbillurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
