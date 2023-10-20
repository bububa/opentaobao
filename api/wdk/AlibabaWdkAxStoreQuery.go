package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkAxStoreQuery 翱象经营店查询接口
// alibaba.wdk.ax.store.query
//
// 翱象经营店查询接口
func AlibabaWdkAxStoreQuery(clt *core.SDKClient, req *wdk.AlibabaWdkAxStoreQueryAPIRequest, resp *wdk.AlibabaWdkAxStoreQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
