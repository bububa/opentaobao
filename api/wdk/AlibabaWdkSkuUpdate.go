package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskuupdate 更新商品
// alibaba.wdk.sku.update
//
// 开放商品更新接口
func Alibabawdkskuupdate(clt *core.SDKClient, req *wdk.AlibabawdkskuupdateAPIRequest, session string) (*wdk.AlibabawdkskuupdateAPIResponse, error) {
	var resp wdk.AlibabawdkskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
