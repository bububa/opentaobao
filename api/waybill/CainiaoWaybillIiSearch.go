package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

/* CainiaoWaybillIiSearch
查询面单服务订购及面单使用情况
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况 */
func CainiaoWaybillIiSearch(clt *core.SDKClient, req *waybill.CainiaoWaybillIiSearchAPIRequest, session string) (*waybill.CainiaoWaybillIiSearchAPIResponse, error) {
	var resp waybill.CainiaoWaybillIiSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
