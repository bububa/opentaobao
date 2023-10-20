package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// Taobaojushitajmsuseradd 添加ONS消息同步用户
// taobao.jushita.jms.user.add
//
// 添加ONS消息同步用户
func Taobaojushitajmsuseradd(clt *core.SDKClient, req *jms.TaobaojushitajmsuseraddAPIRequest, session string) (*jms.TaobaojushitajmsuseraddAPIResponse, error) {
	var resp jms.TaobaojushitajmsuseraddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
