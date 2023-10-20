package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketsend 商家电子凭证发码成功回调接口
// taobao.vmarket.eticket.send
//
// 外部商家成功发码回调接口
func Taobaovmarketeticketsend(clt *core.SDKClient, req *eticket.TaobaovmarketeticketsendAPIRequest, session string) (*eticket.TaobaovmarketeticketsendAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
