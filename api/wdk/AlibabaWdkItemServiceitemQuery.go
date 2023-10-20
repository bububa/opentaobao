package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkitemserviceitemquery 查询服务商品
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
func Alibabawdkitemserviceitemquery(clt *core.SDKClient, req *wdk.AlibabawdkitemserviceitemqueryAPIRequest, session string) (*wdk.AlibabawdkitemserviceitemqueryAPIResponse, error) {
	var resp wdk.AlibabawdkitemserviceitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
