package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdGroupCreateForbiddenProduct
创建屏蔽品
alibaba.scbp.ad.group.create.forbidden.product

创建屏蔽品 */
func AlibabaScbpAdGroupCreateForbiddenProduct(clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCreateForbiddenProductAPIRequest, session string) (*scbp.AlibabaScbpAdGroupCreateForbiddenProductAPIResponse, error) {
	var resp scbp.AlibabaScbpAdGroupCreateForbiddenProductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
