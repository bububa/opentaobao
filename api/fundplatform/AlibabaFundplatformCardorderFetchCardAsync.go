package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardorderfetchcardasync 异步批量生成储值卡
// alibaba.fundplatform.cardorder.fetch.card.async
//
// 外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
func Alibabafundplatformcardorderfetchcardasync(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardorderfetchcardasyncAPIRequest, session string) (*fundplatform.AlibabafundplatformcardorderfetchcardasyncAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardorderfetchcardasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
