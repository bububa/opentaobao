package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Cainiaocbossworkplatformlogisticsiscainiaoorder 根据交易单号判断是否为菜鸟发货订单
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
func Cainiaocbossworkplatformlogisticsiscainiaoorder(clt *core.SDKClient, req *logistic.CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest, session string) (*logistic.CainiaocbossworkplatformlogisticsiscainiaoorderAPIResponse, error) {
	var resp logistic.CainiaocbossworkplatformlogisticsiscainiaoorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
