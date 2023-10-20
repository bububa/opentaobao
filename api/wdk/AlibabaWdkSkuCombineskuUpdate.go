package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskucombineskuupdate 组合商品更新接口
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
func Alibabawdkskucombineskuupdate(clt *core.SDKClient, req *wdk.AlibabawdkskucombineskuupdateAPIRequest, session string) (*wdk.AlibabawdkskucombineskuupdateAPIResponse, error) {
	var resp wdk.AlibabawdkskucombineskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
