package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeentrustsellingquery 委托信息查询接口
// alibaba.alihouse.existinghome.entrustselling.query
//
// 管家状态及房源信息接口
func Alibabaalihouseexistinghomeentrustsellingquery(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeentrustsellingqueryAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeentrustsellingqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeentrustsellingqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
