package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcApContracturlGet 按用户获取支付宝代扣协议链接地址
// taobao.oc.ap.contracturl.get
//
// 按用户获取支付宝代扣协议链接地址
func TaobaoOcApContracturlGet(clt *core.SDKClient, req *jst.TaobaoOcApContracturlGetAPIRequest, session string) (*jst.TaobaoOcApContracturlGetAPIResponse, error) {
	var resp jst.TaobaoOcApContracturlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
