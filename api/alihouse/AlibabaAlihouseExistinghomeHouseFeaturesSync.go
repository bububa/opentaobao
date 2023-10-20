package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousefeaturessync 房源标准打标数据同步
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
func Alibabaalihouseexistinghomehousefeaturessync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousefeaturessyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousefeaturessyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousefeaturessyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
