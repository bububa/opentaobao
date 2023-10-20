package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityIdentityStore 用户信息存储
// alibaba.jhs.community.identity.store
//
// 用户信息存储
func AlibabaJhsCommunityIdentityStore(clt *core.SDKClient, req *ju.AlibabaJhsCommunityIdentityStoreAPIRequest, session string) (*ju.AlibabaJhsCommunityIdentityStoreAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityIdentityStoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
