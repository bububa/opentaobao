package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// Taobaojushitajmsgroupget 查询ONS分组
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
func Taobaojushitajmsgroupget(clt *core.SDKClient, req *jms.TaobaojushitajmsgroupgetAPIRequest, session string) (*jms.TaobaojushitajmsgroupgetAPIResponse, error) {
	var resp jms.TaobaojushitajmsgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
