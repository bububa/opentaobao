package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupDeleteForbiddenProduct 删除屏蔽品
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
func AlibabaScbpAdGroupDeleteForbiddenProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest, resp *scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
