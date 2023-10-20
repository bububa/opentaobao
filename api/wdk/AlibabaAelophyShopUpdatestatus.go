package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyshopupdatestatus 更新渠道店营业状态
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
func Alibabaaelophyshopupdatestatus(clt *core.SDKClient, req *wdk.AlibabaaelophyshopupdatestatusAPIRequest, session string) (*wdk.AlibabaaelophyshopupdatestatusAPIResponse, error) {
	var resp wdk.AlibabaaelophyshopupdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
