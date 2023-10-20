package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharityuserbindgeturi 获取用户绑定uri
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
func Alibabacharityuserbindgeturi(clt *core.SDKClient, req *charity.AlibabacharityuserbindgeturiAPIRequest, session string) (*charity.AlibabacharityuserbindgeturiAPIResponse, error) {
	var resp charity.AlibabacharityuserbindgeturiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
