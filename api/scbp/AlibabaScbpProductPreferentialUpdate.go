package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpProductPreferentialUpdate 设置P4P产品优先推广状态
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
func AlibabaScbpProductPreferentialUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpProductPreferentialUpdateAPIRequest, resp *scbp.AlibabaScbpProductPreferentialUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
