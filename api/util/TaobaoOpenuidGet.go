package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoopenuidget 获取授权账号对应的OpenUid
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
func Taobaoopenuidget(clt *core.SDKClient, req *util.TaobaoopenuidgetAPIRequest, session string) (*util.TaobaoopenuidgetAPIResponse, error) {
	var resp util.TaobaoopenuidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
