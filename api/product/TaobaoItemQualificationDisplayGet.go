package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoitemqualificationdisplayget 资质采集配置异步获取接口
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
func Taobaoitemqualificationdisplayget(clt *core.SDKClient, req *product.TaobaoitemqualificationdisplaygetAPIRequest, session string) (*product.TaobaoitemqualificationdisplaygetAPIResponse, error) {
	var resp product.TaobaoitemqualificationdisplaygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
