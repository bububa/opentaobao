package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// Cainiaoguoguocpbackupassigncourierbyid 根据菜鸟账号ID指派小件员
// cainiao.guoguo.cp.backup.assigncourierbyid
//
// 根据菜鸟账号ID指派小件员
func Cainiaoguoguocpbackupassigncourierbyid(clt *core.SDKClient, req *guoguo.CainiaoguoguocpbackupassigncourierbyidAPIRequest, session string) (*guoguo.CainiaoguoguocpbackupassigncourierbyidAPIResponse, error) {
	var resp guoguo.CainiaoguoguocpbackupassigncourierbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
