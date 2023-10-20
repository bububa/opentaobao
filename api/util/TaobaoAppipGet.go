package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoappipget 获取ISV发起请求服务器IP
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
func Taobaoappipget(clt *core.SDKClient, req *util.TaobaoappipgetAPIRequest, session string) (*util.TaobaoappipgetAPIResponse, error) {
	var resp util.TaobaoappipgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
