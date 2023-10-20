package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupFindForbiddenProduct 查询屏蔽品
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
func AlibabaScbpAdGroupFindForbiddenProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupFindForbiddenProductAPIRequest, resp *scbp.AlibabaScbpAdGroupFindForbiddenProductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
