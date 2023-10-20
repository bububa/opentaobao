package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// Taobaojushitajmstopicsget 根据用户nick获取开通的消息列表
// taobao.jushita.jms.topics.get
//
// 根据用户nick获取开通的消息列表
func Taobaojushitajmstopicsget(clt *core.SDKClient, req *jms.TaobaojushitajmstopicsgetAPIRequest, session string) (*jms.TaobaojushitajmstopicsgetAPIResponse, error) {
	var resp jms.TaobaojushitajmstopicsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
