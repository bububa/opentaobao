package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupDeleteForbiddenProduct 删除屏蔽品
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
func AlibabaScbpAdGroupDeleteForbiddenProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest, session string) (*scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
