package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeentrustsellingupdate 管家状态及房源信息接口
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
func Alibabaalihouseexistinghomeentrustsellingupdate(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeentrustsellingupdateAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeentrustsellingupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeentrustsellingupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
