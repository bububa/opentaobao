package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaogloballogisticordercreate 创建物流订单
// cainiao.global.logistic.order.create
//
// 创建物流订单
func Cainiaogloballogisticordercreate(clt *core.SDKClient, req *cainiaohandover.CainiaogloballogisticordercreateAPIRequest, session string) (*cainiaohandover.CainiaogloballogisticordercreateAPIResponse, error) {
	var resp cainiaohandover.CainiaogloballogisticordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
