package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// Cainiaoguoguocpbackupassigncourier CP兜底后指定接单的小件员
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
func Cainiaoguoguocpbackupassigncourier(clt *core.SDKClient, req *guoguo.CainiaoguoguocpbackupassigncourierAPIRequest, session string) (*guoguo.CainiaoguoguocpbackupassigncourierAPIResponse, error) {
	var resp guoguo.CainiaoguoguocpbackupassigncourierAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
