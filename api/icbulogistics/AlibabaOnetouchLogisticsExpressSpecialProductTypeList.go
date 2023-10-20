package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressSpecialProductTypeList 获取商品类型配置项
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
func AlibabaOnetouchLogisticsExpressSpecialProductTypeList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
