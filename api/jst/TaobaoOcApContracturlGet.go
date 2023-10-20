package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaoocapcontracturlget 按用户获取支付宝代扣协议链接地址
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
func Taobaoocapcontracturlget(clt *core.SDKClient, req *jst.TaobaoocapcontracturlgetAPIRequest, session string) (*jst.TaobaoocapcontracturlgetAPIResponse, error) {
	var resp jst.TaobaoocapcontracturlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
