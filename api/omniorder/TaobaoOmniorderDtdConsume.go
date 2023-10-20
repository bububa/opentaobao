package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderDtdConsume 门店自送对码进行核销
// taobao.omniorder.dtd.consume
//
// 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
func TaobaoOmniorderDtdConsume(clt *core.SDKClient, req *omniorder.TaobaoOmniorderDtdConsumeAPIRequest, session string) (*omniorder.TaobaoOmniorderDtdConsumeAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderDtdConsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
