package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizapactivate 激活AP价签通讯模块
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
func Taobaousceslbizapactivate(clt *core.SDKClient, req *uscesl.TaobaousceslbizapactivateAPIRequest, session string) (*uscesl.TaobaousceslbizapactivateAPIResponse, error) {
	var resp uscesl.TaobaousceslbizapactivateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
