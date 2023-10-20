package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserBindGeturi 获取用户绑定uri
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
func AlibabaCharityUserBindGeturi(clt *core.SDKClient, req *charity.AlibabaCharityUserBindGeturiAPIRequest, session string) (*charity.AlibabaCharityUserBindGeturiAPIResponse, error) {
	var resp charity.AlibabaCharityUserBindGeturiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
