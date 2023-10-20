package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautotradecareticketavailablecheck 天猫汽车电子凭证核销前校验
// tmall.aliauto.trade.car.eticket.available.check
//
// 天猫汽车核销码可用性校验
func Tmallaliautotradecareticketavailablecheck(clt *core.SDKClient, req *tmallcar.TmallaliautotradecareticketavailablecheckAPIRequest, session string) (*tmallcar.TmallaliautotradecareticketavailablecheckAPIResponse, error) {
	var resp tmallcar.TmallaliautotradecareticketavailablecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
