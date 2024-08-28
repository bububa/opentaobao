package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectKanameQuery 查询KA楼盘名称
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
func AlibabaAlihouseNewhomeProjectKanameQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
