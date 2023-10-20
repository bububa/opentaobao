package cntms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cntms"
)

// Cainiaocntmslogisticsorderconsign 菜鸟配商家仓库发货
// cainiao.cntms.logistics.order.consign
//
// 商家包装打印面单结束后，通知菜鸟包裹要发货
func Cainiaocntmslogisticsorderconsign(clt *core.SDKClient, req *cntms.CainiaocntmslogisticsorderconsignAPIRequest, session string) (*cntms.CainiaocntmslogisticsorderconsignAPIResponse, error) {
	var resp cntms.CainiaocntmslogisticsorderconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
