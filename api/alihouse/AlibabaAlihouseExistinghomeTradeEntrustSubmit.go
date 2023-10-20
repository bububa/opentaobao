package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghometradeentrustsubmit 交易委托信息更新接口
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
func Alibabaalihouseexistinghometradeentrustsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghometradeentrustsubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghometradeentrustsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghometradeentrustsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
