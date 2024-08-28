package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemOuteridUpdate 天猫商品/SKU商家编码更新接口
// tmall.item.outerid.update
//
// 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
func TmallItemOuteridUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemOuteridUpdateAPIRequest, resp *tbitem.TmallItemOuteridUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
