package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// Taobaojushitajmsuserget 查询某个用户是否同步消息
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
func Taobaojushitajmsuserget(clt *core.SDKClient, req *jms.TaobaojushitajmsusergetAPIRequest, session string) (*jms.TaobaojushitajmsusergetAPIResponse, error) {
	var resp jms.TaobaojushitajmsusergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
