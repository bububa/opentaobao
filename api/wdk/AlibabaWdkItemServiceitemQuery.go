package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemServiceitemQuery 查询服务商品
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
func AlibabaWdkItemServiceitemQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkItemServiceitemQueryAPIRequest, resp *wdk.AlibabaWdkItemServiceitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
