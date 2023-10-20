package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharityuserexternalauthcancel 取消外部授权
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
func Alibabacharityuserexternalauthcancel(clt *core.SDKClient, req *charity.AlibabacharityuserexternalauthcancelAPIRequest, session string) (*charity.AlibabacharityuserexternalauthcancelAPIResponse, error) {
	var resp charity.AlibabacharityuserexternalauthcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
