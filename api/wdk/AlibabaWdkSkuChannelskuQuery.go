package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuchannelskuquery 查询渠道商品
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
func Alibabawdkskuchannelskuquery(clt *core.SDKClient, req *wdk.AlibabawdkskuchannelskuqueryAPIRequest, session string) (*wdk.AlibabawdkskuchannelskuqueryAPIResponse, error) {
	var resp wdk.AlibabawdkskuchannelskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
