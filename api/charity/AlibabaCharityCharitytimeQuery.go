package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimequery 查询公益3小时公益时汇总
// alibaba.charity.charitytime.query
//
// 查询公益3小时公益时汇总
func Alibabacharitycharitytimequery(clt *core.SDKClient, req *charity.AlibabacharitycharitytimequeryAPIRequest, session string) (*charity.AlibabacharitycharitytimequeryAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
