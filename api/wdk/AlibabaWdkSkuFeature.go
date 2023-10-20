package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkskufeature 商品标记接口
// alibaba.wdk.sku.feature
//
// 给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
func Alibabawdkskufeature(clt *core.SDKClient, req *wdk.AlibabawdkskufeatureAPIRequest, session string) (*wdk.AlibabawdkskufeatureAPIResponse, error) {
	var resp wdk.AlibabawdkskufeatureAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
