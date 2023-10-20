package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpressspecialproducttypelist 获取商品类型配置项
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
func Alibabaonetouchlogisticsexpressspecialproducttypelist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpressspecialproducttypelistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpressspecialproducttypelistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpressspecialproducttypelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
