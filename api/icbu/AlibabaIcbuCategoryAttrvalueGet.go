package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryAttrvalueGet 属性值获取
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
func AlibabaIcbuCategoryAttrvalueGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryAttrvalueGetAPIRequest, resp *icbu.AlibabaIcbuCategoryAttrvalueGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
