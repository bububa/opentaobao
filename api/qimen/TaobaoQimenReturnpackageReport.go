package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenreturnpackagereport 退货包裹状态通知接口
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
func Taobaoqimenreturnpackagereport(clt *core.SDKClient, req *qimen.TaobaoqimenreturnpackagereportAPIRequest, session string) (*qimen.TaobaoqimenreturnpackagereportAPIResponse, error) {
	var resp qimen.TaobaoqimenreturnpackagereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
