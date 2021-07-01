package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

/* AliexpressLogisticsBuyerFreightCalculate
提供给买家使用的运费计算接口
aliexpress.logistics.buyer.freight.calculate

提供给买家使用的运费计算接口 */
func AliexpressLogisticsBuyerFreightCalculate(clt *core.SDKClient, req *aedropshiper.AliexpressLogisticsBuyerFreightCalculateAPIRequest, session string) (*aedropshiper.AliexpressLogisticsBuyerFreightCalculateAPIResponse, error) {
	var resp aedropshiper.AliexpressLogisticsBuyerFreightCalculateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
