package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghousehousebase （租房）同步房屋信息
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
func Alibabaalihouseexistinghousehousebase(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghousehousebaseAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghousehousebaseAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghousehousebaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
