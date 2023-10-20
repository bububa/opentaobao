package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsaddressremove 删除卖家地址库
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
func Taobaologisticsaddressremove(clt *core.SDKClient, req *logistic.TaobaologisticsaddressremoveAPIRequest, session string) (*logistic.TaobaologisticsaddressremoveAPIResponse, error) {
	var resp logistic.TaobaologisticsaddressremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
