package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaomodifyorderordercheck 自助改单服务发货订单校验
// taobao.modifyorder.order.check
//
// 自助改单服务发货后订单回传接口
func Taobaomodifyorderordercheck(clt *core.SDKClient, req *jst.TaobaomodifyorderordercheckAPIRequest, session string) (*jst.TaobaomodifyorderordercheckAPIResponse, error) {
	var resp jst.TaobaomodifyorderordercheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
