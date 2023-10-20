package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaoendpointlockertopstationaddorupdate 增加更新代收点
// cainiao.endpoint.locker.top.station.addorupdate
//
// 新增或者修改代收点相关信息
func Cainiaoendpointlockertopstationaddorupdate(clt *core.SDKClient, req *cainiaolocker.CainiaoendpointlockertopstationaddorupdateAPIRequest, session string) (*cainiaolocker.CainiaoendpointlockertopstationaddorupdateAPIResponse, error) {
	var resp cainiaolocker.CainiaoendpointlockertopstationaddorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
