package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpEffectProductSingleGet
单个产品的报表
alibaba.scbp.effect.product.single.get

单个产品的报表 */
func AlibabaScbpEffectProductSingleGet(clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductSingleGetAPIRequest, session string) (*scbp.AlibabaScbpEffectProductSingleGetAPIResponse, error) {
	var resp scbp.AlibabaScbpEffectProductSingleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
