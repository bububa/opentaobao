package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreUpdate 翱翔经营店更新接口
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
func AlibabaWdkAxStoreUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreUpdateAPIRequest, resp *wdk.AlibabaWdkAxStoreUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
