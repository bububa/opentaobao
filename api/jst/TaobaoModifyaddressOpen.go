package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaomodifyaddressopen 淘宝自助修改地址服务开通
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
func Taobaomodifyaddressopen(clt *core.SDKClient, req *jst.TaobaomodifyaddressopenAPIRequest, session string) (*jst.TaobaomodifyaddressopenAPIResponse, error) {
	var resp jst.TaobaomodifyaddressopenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
