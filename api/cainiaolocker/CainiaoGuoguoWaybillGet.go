package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaoguoguowaybillget 菜鸟裹裹商家寄件取号接口
// cainiao.guoguo.waybill.get
//
// 菜鸟裹裹商家寄件取号接口
func Cainiaoguoguowaybillget(clt *core.SDKClient, req *cainiaolocker.CainiaoguoguowaybillgetAPIRequest, session string) (*cainiaolocker.CainiaoguoguowaybillgetAPIResponse, error) {
	var resp cainiaolocker.CainiaoguoguowaybillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
