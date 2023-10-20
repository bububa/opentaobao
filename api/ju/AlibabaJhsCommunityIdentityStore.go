package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunityidentitystore 用户信息存储
// alibaba.jhs.community.identity.store
//
// 用户信息存储
func Alibabajhscommunityidentitystore(clt *core.SDKClient, req *ju.AlibabajhscommunityidentitystoreAPIRequest, session string) (*ju.AlibabajhscommunityidentitystoreAPIResponse, error) {
	var resp ju.AlibabajhscommunityidentitystoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
