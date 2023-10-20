package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyshopupdateinfo 更新渠道店基础信息
// alibaba.aelophy.shop.updateinfo
//
// 更新渠道店基础信息
func Alibabaaelophyshopupdateinfo(clt *core.SDKClient, req *wdk.AlibabaaelophyshopupdateinfoAPIRequest, session string) (*wdk.AlibabaaelophyshopupdateinfoAPIResponse, error) {
	var resp wdk.AlibabaaelophyshopupdateinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
