package c2m

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

// Taobaosebporganizationgetorderinfo 淘小铺机构订单信息
// taobao.sebp.organization.getorderinfo
//
// 淘小铺合作机构获取机构订单信息，用于机构结算使用
func Taobaosebporganizationgetorderinfo(clt *core.SDKClient, req *c2m.TaobaosebporganizationgetorderinfoAPIRequest, session string) (*c2m.TaobaosebporganizationgetorderinfoAPIResponse, error) {
	var resp c2m.TaobaosebporganizationgetorderinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
