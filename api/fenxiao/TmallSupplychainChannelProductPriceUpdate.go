package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TmallSupplychainChannelProductPriceUpdate
渠道价格更新接口
tmall.supplychain.channel.product.price.update

更新渠道产品价格 */
func TmallSupplychainChannelProductPriceUpdate(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductPriceUpdateAPIRequest, session string) (*fenxiao.TmallSupplychainChannelProductPriceUpdateAPIResponse, error) {
	var resp fenxiao.TmallSupplychainChannelProductPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
