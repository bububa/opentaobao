package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqRecommend rfq推荐
// alibaba.icbu.rfq.recommend
//
// rfq推荐
func AlibabaIcbuRfqRecommend(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqRecommendAPIRequest, resp *icburfq.AlibabaIcbuRfqRecommendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
