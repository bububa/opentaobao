package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemServiceitemQuery 查询服务商品
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
func AlibabaWdkItemServiceitemQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemServiceitemQueryAPIRequest, resp *wdk.AlibabaWdkItemServiceitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
