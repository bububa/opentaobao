package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabaomnisaasordercreate 订单创建接口
// alibaba.omni.saas.order.create
//
// 服务商利用现有的saas系统和阿里完成交易系统的对接
func Alibabaomnisaasordercreate(clt *core.SDKClient, req *trade.AlibabaomnisaasordercreateAPIRequest, session string) (*trade.AlibabaomnisaasordercreateAPIResponse, error) {
	var resp trade.AlibabaomnisaasordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
