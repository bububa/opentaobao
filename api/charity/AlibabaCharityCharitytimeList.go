package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimelist 授权用户的公益时记录查询
// alibaba.charity.charitytime.list
//
// 查询授权用户的公益时记录
func Alibabacharitycharitytimelist(clt *core.SDKClient, req *charity.AlibabacharitycharitytimelistAPIRequest, session string) (*charity.AlibabacharitycharitytimelistAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
