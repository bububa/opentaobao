package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkshopquery 门店查询接口
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
func Alibabawdkshopquery(clt *core.SDKClient, req *wdk.AlibabawdkshopqueryAPIRequest, session string) (*wdk.AlibabawdkshopqueryAPIResponse, error) {
	var resp wdk.AlibabawdkshopqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
