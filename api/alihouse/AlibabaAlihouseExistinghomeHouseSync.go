package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousesync 房源信息同步
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
func Alibabaalihouseexistinghomehousesync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousesyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousesyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
