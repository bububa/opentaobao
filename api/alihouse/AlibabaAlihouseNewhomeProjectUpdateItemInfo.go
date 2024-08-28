package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectUpdateItemInfo 更新楼盘商品信息
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
func AlibabaAlihouseNewhomeProjectUpdateItemInfo(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
