package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectQuery 查询楼盘相关信息
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
func AlibabaAlihouseNewhomeProjectQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectQueryAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
