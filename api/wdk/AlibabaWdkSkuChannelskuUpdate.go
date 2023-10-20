package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuchannelskuupdate 更新渠道商品
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
func Alibabawdkskuchannelskuupdate(clt *core.SDKClient, req *wdk.AlibabawdkskuchannelskuupdateAPIRequest, session string) (*wdk.AlibabawdkskuchannelskuupdateAPIResponse, error) {
	var resp wdk.AlibabawdkskuchannelskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
