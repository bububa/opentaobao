package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Cainiaocntecsupplierorderservice 供货商查询订单数据接口
// cainiao.cntec.supplier.order.service
//
// 提供给供货商查询订单信息的接口，返回给供货商的订单数据已经是脱敏精简后的，比如订单ID用户ID已经用md5加密，用户昵称已经脱敏，商品信息本身是供货商提供的。
// 数据查询的范围只和供货商的身份有关系，比如大润发的用户只能查询大润发的订单，而且会校验身份和颁发的appkey之间的关系，并且目前对接的只有一个供货商
func Cainiaocntecsupplierorderservice(clt *core.SDKClient, req *trade.CainiaocntecsupplierorderserviceAPIRequest, session string) (*trade.CainiaocntecsupplierorderserviceAPIResponse, error) {
	var resp trade.CainiaocntecsupplierorderserviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
