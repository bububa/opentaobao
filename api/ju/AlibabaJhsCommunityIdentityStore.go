package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityIdentityStore 用户信息存储
// alibaba.jhs.community.identity.store
//
// 用户信息存储
func AlibabaJhsCommunityIdentityStore(clt *core.SDKClient, req *ju.AlibabaJhsCommunityIdentityStoreAPIRequest, resp *ju.AlibabaJhsCommunityIdentityStoreAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
