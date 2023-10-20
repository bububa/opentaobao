package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserBindGeturi 获取用户绑定uri
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
func AlibabaCharityUserBindGeturi(clt *core.SDKClient, req *charity.AlibabaCharityUserBindGeturiAPIRequest, resp *charity.AlibabaCharityUserBindGeturiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
