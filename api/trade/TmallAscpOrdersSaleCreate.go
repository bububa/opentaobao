package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Tmallascporderssalecreate ASCP渠道中心销售单创建接口
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
func Tmallascporderssalecreate(clt *core.SDKClient, req *trade.TmallascporderssalecreateAPIRequest, session string) (*trade.TmallascporderssalecreateAPIResponse, error) {
	var resp trade.TmallascporderssalecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
