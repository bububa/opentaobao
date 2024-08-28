package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreCreate 翱象经营店创建接口
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
func AlibabaWdkAxStoreCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreCreateAPIRequest, resp *wdk.AlibabaWdkAxStoreCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
