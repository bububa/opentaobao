package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemQualificationDisplayGet 资质采集配置异步获取接口
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
func TaobaoItemQualificationDisplayGet(ctx context.Context, clt *core.SDKClient, req *product.TaobaoItemQualificationDisplayGetAPIRequest, resp *product.TaobaoItemQualificationDisplayGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
