package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautotraderestpayfeemodify 分阶段订单尾款改价
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
func Tmallaliautotraderestpayfeemodify(clt *core.SDKClient, req *tmallcar.TmallaliautotraderestpayfeemodifyAPIRequest, session string) (*tmallcar.TmallaliautotraderestpayfeemodifyAPIResponse, error) {
	var resp tmallcar.TmallaliautotraderestpayfeemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
