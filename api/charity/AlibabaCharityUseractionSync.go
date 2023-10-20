package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacharityuseractionsync 用户公益行为同步
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
func Alibabacharityuseractionsync(clt *core.SDKClient, req *charity.AlibabacharityuseractionsyncAPIRequest, session string) (*charity.AlibabacharityuseractionsyncAPIResponse, error) {
	var resp charity.AlibabacharityuseractionsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
