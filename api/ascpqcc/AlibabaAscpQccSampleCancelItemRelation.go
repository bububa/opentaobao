package ascpqcc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpqcc"
)

// AlibabaAscpQccSampleCancelItemRelation 魅力惠样品解除父子商品关系
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
func AlibabaAscpQccSampleCancelItemRelation(clt *core.SDKClient, req *ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIRequest, resp *ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
