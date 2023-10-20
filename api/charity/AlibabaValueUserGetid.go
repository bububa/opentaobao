package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaValueUserGetid 获取用户userId
// alibaba.value.user.getid
//
// 获取用户userId
func AlibabaValueUserGetid(clt *core.SDKClient, req *charity.AlibabaValueUserGetidAPIRequest, session string) (*charity.AlibabaValueUserGetidAPIResponse, error) {
	var resp charity.AlibabaValueUserGetidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
