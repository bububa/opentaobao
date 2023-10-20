package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabuseropenuidget access token 获取精灵用户 id
// alibaba.ailab.user.open.uid.get
//
// access token 获取精灵用户 id
func Alibabaailabuseropenuidget(clt *core.SDKClient, req *alilabs.AlibabaailabuseropenuidgetAPIRequest, session string) (*alilabs.AlibabaailabuseropenuidgetAPIResponse, error) {
	var resp alilabs.AlibabaailabuseropenuidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
