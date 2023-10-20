package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreUpdate 翱翔经营店更新接口
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
func AlibabaWdkAxStoreUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreUpdateAPIRequest, resp *wdk.AlibabaWdkAxStoreUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
