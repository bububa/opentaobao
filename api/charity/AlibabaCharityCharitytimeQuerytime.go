package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharitycharitytimequerytime 外部查询公益时
// alibaba.charity.charitytime.querytime
//
// 外部查询公益时
func Alibabacharitycharitytimequerytime(clt *core.SDKClient, req *charity.AlibabacharitycharitytimequerytimeAPIRequest, session string) (*charity.AlibabacharitycharitytimequerytimeAPIResponse, error) {
	var resp charity.AlibabacharitycharitytimequerytimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
