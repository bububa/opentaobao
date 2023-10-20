package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketcodesget 电子凭证码列表查询
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
func Taobaovmarketeticketcodesget(clt *core.SDKClient, req *eticket.TaobaovmarketeticketcodesgetAPIRequest, session string) (*eticket.TaobaovmarketeticketcodesgetAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketcodesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
