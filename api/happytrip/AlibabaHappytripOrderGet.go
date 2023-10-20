package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriporderget 获取欢行统一订单模型
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
func Alibabahappytriporderget(clt *core.SDKClient, req *happytrip.AlibabahappytripordergetAPIRequest, session string) (*happytrip.AlibabahappytripordergetAPIResponse, error) {
	var resp happytrip.AlibabahappytripordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
