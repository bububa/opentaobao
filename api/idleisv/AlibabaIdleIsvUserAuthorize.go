package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvuserauthorize 用户授权接口
// alibaba.idle.isv.user.authorize
//
// 用户授权接口
// 接口调用相关参考文档
// https://www.yuque.com/docs/share/9cd991b7-e3a3-40b6-948c-1835422d0164?# 《闲鱼优品2.0API接入说明》
func Alibabaidleisvuserauthorize(clt *core.SDKClient, req *idleisv.AlibabaidleisvuserauthorizeAPIRequest, session string) (*idleisv.AlibabaidleisvuserauthorizeAPIResponse, error) {
	var resp idleisv.AlibabaidleisvuserauthorizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
