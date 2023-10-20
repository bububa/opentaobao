package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// Cainiaoguoguocpnborderfrontrupdateuser 小件员信息变更
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
func Cainiaoguoguocpnborderfrontrupdateuser(clt *core.SDKClient, req *guoguo.CainiaoguoguocpnborderfrontrupdateuserAPIRequest, session string) (*guoguo.CainiaoguoguocpnborderfrontrupdateuserAPIResponse, error) {
	var resp guoguo.CainiaoguoguocpnborderfrontrupdateuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
