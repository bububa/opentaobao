package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticscompaniesget 查询物流公司信息
// taobao.logistics.companies.get
//
// 查询淘宝网合作的物流公司信息，用于发货接口。
func Taobaologisticscompaniesget(clt *core.SDKClient, req *tblogistics.TaobaologisticscompaniesgetAPIRequest, session string) (*tblogistics.TaobaologisticscompaniesgetAPIResponse, error) {
	var resp tblogistics.TaobaologisticscompaniesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
