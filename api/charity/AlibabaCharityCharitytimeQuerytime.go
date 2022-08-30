package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeQuerytime 外部查询公益时
// alibaba.charity.charitytime.querytime
//
// 外部查询公益时
func AlibabaCharityCharitytimeQuerytime(clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeQuerytimeAPIRequest, session string) (*charity.AlibabaCharityCharitytimeQuerytimeAPIResponse, error) {
	var resp charity.AlibabaCharityCharitytimeQuerytimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
