package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojecttradeorder 旺铺楼盘和交易商品排序
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
func Alibabaalihousenewhomeprojecttradeorder(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojecttradeorderAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojecttradeorderAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojecttradeorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
