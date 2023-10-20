package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitybindcancel 取消用户绑定
// alibaba.charity.bind.cancel
//
// 取消用户绑定
func Alibabacharitybindcancel(clt *core.SDKClient, req *charity.AlibabacharitybindcancelAPIRequest, session string) (*charity.AlibabacharitybindcancelAPIResponse, error) {
	var resp charity.AlibabacharitybindcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
