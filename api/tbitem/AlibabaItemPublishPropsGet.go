package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishPropsGet 商品级联属性信息获取
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
func AlibabaItemPublishPropsGet(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemPublishPropsGetAPIRequest, resp *tbitem.AlibabaItemPublishPropsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
