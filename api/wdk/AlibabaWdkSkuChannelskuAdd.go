package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuchannelskuadd 新增渠道商品
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
func Alibabawdkskuchannelskuadd(clt *core.SDKClient, req *wdk.AlibabawdkskuchannelskuaddAPIRequest, session string) (*wdk.AlibabawdkskuchannelskuaddAPIResponse, error) {
	var resp wdk.AlibabawdkskuchannelskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
