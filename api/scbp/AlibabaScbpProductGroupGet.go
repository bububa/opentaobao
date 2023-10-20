package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpProductGroupGet 查询指定产品分组的下一层子分组
// alibaba.scbp.product.group.get
//
// 查询指定产品分组的下一层子分组
func AlibabaScbpProductGroupGet(clt *core.SDKClient, req *scbp.AlibabaScbpProductGroupGetAPIRequest, resp *scbp.AlibabaScbpProductGroupGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
