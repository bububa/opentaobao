package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreUpdate 翱翔经营店更新接口
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
func AlibabaWdkAxStoreUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreUpdateAPIRequest, session string) (*wdk.AlibabaWdkAxStoreUpdateAPIResponse, error) {
	var resp wdk.AlibabaWdkAxStoreUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
