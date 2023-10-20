package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsaddressadd 卖家地址库新增接口
// taobao.logistics.address.add
//
// 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库
func Taobaologisticsaddressadd(clt *core.SDKClient, req *logistic.TaobaologisticsaddressaddAPIRequest, session string) (*logistic.TaobaologisticsaddressaddAPIResponse, error) {
	var resp logistic.TaobaologisticsaddressaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
