package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseCategoryControlQuery 类目权限查询
// alibaba.alihouse.category.control.query
//
// 类目权限查询
func AlibabaAlihouseCategoryControlQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseCategoryControlQueryAPIRequest, resp *alihouse.AlibabaAlihouseCategoryControlQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
