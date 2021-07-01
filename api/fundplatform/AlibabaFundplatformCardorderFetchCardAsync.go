package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

/* AlibabaFundplatformCardorderFetchCardAsync
异步批量生成储值卡
alibaba.fundplatform.cardorder.fetch.card.async

外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功 */
func AlibabaFundplatformCardorderFetchCardAsync(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderFetchCardAsyncAPIRequest, session string) (*fundplatform.AlibabaFundplatformCardorderFetchCardAsyncAPIResponse, error) {
	var resp fundplatform.AlibabaFundplatformCardorderFetchCardAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
