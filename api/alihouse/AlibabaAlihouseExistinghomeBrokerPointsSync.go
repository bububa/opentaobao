package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrokerpointssync 经纪人积分同步
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
func Alibabaalihouseexistinghomebrokerpointssync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrokerpointssyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrokerpointssyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrokerpointssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
