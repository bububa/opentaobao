package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Cainiaocrmomsrulesync 商家ERP订单处理规则同步
// cainiao.crm.oms.rule.sync
//
// 将商家ERP订单处理规则同步到菜鸟CRM系统
func Cainiaocrmomsrulesync(clt *core.SDKClient, req *wms.CainiaocrmomsrulesyncAPIRequest, session string) (*wms.CainiaocrmomsrulesyncAPIResponse, error) {
	var resp wms.CainiaocrmomsrulesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
