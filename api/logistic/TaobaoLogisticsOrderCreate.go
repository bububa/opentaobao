package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsOrderCreate 创建物流订单
// taobao.logistics.order.create
//
// 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
func TaobaoLogisticsOrderCreate(clt *core.SDKClient, req *logistic.TaobaoLogisticsOrderCreateAPIRequest, session string) (*logistic.TaobaoLogisticsOrderCreateAPIResponse, error) {
	var resp logistic.TaobaoLogisticsOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
