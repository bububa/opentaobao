package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyshopupdaterange 更新渠道店销售范围
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
func Alibabaaelophyshopupdaterange(clt *core.SDKClient, req *wdk.AlibabaaelophyshopupdaterangeAPIRequest, session string) (*wdk.AlibabaaelophyshopupdaterangeAPIResponse, error) {
	var resp wdk.AlibabaaelophyshopupdaterangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
