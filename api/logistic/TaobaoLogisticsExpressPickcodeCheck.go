package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpresspickcodecheck 快递公司取货码校验
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
func Taobaologisticsexpresspickcodecheck(clt *core.SDKClient, req *logistic.TaobaologisticsexpresspickcodecheckAPIRequest, session string) (*logistic.TaobaologisticsexpresspickcodecheckAPIResponse, error) {
	var resp logistic.TaobaologisticsexpresspickcodecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
