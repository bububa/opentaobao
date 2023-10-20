package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsaddressmodify 卖家地址库修改
// taobao.logistics.address.modify
//
// 卖家地址库修改
func Taobaologisticsaddressmodify(clt *core.SDKClient, req *logistic.TaobaologisticsaddressmodifyAPIRequest, session string) (*logistic.TaobaologisticsaddressmodifyAPIResponse, error) {
	var resp logistic.TaobaologisticsaddressmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
