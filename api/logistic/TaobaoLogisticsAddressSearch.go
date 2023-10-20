package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsaddresssearch 查询卖家地址库
// taobao.logistics.address.search
//
// 通过此接口查询卖家地址库，
func Taobaologisticsaddresssearch(clt *core.SDKClient, req *logistic.TaobaologisticsaddresssearchAPIRequest, session string) (*logistic.TaobaologisticsaddresssearchAPIResponse, error) {
	var resp logistic.TaobaologisticsaddresssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
