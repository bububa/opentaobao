package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoServicePackageQuery 预采购服务包查询接口
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
func AlibabaEleFengniaoServicePackageQuery(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoServicePackageQueryAPIRequest, resp *logistic.AlibabaEleFengniaoServicePackageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
