package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrorderlogisinfo 区域零售订单获取取件码
// tmall.nr.order.logis.info
//
// 区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
func Tmallnrorderlogisinfo(clt *core.SDKClient, req *tmallnr.TmallnrorderlogisinfoAPIRequest, session string) (*tmallnr.TmallnrorderlogisinfoAPIResponse, error) {
	var resp tmallnr.TmallnrorderlogisinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
