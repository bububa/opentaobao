package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixorderconfirm 大麦-出票
// alibaba.damai.maitix.order.confirm
//
// 出票
func Alibabadamaimaitixorderconfirm(clt *core.SDKClient, req *maitix.AlibabadamaimaitixorderconfirmAPIRequest, session string) (*maitix.AlibabadamaimaitixorderconfirmAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
