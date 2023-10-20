package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxidriverblacklistremove 移除司机黑名单
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
func Alibabahappytriptaxidriverblacklistremove(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxidriverblacklistremoveAPIRequest, session string) (*happytrip.AlibabahappytriptaxidriverblacklistremoveAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxidriverblacklistremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
