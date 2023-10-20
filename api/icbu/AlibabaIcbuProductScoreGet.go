package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductScoreGet 产品质量分查询
// alibaba.icbu.product.score.get
//
// 产品质量分查询
func AlibabaIcbuProductScoreGet(clt *core.SDKClient, req *icbu.AlibabaIcbuProductScoreGetAPIRequest, resp *icbu.AlibabaIcbuProductScoreGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
