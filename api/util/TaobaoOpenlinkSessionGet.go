package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoopenlinksessionget 获取授权session信息
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
func Taobaoopenlinksessionget(clt *core.SDKClient, req *util.TaobaoopenlinksessiongetAPIRequest, session string) (*util.TaobaoopenlinksessiongetAPIResponse, error) {
	var resp util.TaobaoopenlinksessiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
