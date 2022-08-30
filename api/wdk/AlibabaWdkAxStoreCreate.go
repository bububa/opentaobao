package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreCreate 翱象经营店创建接口
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
func AlibabaWdkAxStoreCreate(clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreCreateAPIRequest, session string) (*wdk.AlibabaWdkAxStoreCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkAxStoreCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
