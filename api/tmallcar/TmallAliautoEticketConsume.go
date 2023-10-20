package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoeticketconsume 天猫汽车二轮车电子凭证核销
// tmall.aliauto.eticket.consume
//
// 天猫汽车二轮车行业门店电子凭证核销
func Tmallaliautoeticketconsume(clt *core.SDKClient, req *tmallcar.TmallaliautoeticketconsumeAPIRequest, session string) (*tmallcar.TmallaliautoeticketconsumeAPIResponse, error) {
	var resp tmallcar.TmallaliautoeticketconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
