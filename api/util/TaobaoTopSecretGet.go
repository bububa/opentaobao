package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopsecretget 获取TOP通道解密秘钥
// taobao.top.secret.get
//
// top sdk通过api获取对应解密秘钥
func Taobaotopsecretget(clt *core.SDKClient, req *util.TaobaotopsecretgetAPIRequest, session string) (*util.TaobaotopsecretgetAPIResponse, error) {
	var resp util.TaobaotopsecretgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
