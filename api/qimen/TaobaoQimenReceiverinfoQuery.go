package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenreceiverinfoquery OAID 收件人信息解密接口
// taobao.qimen.receiverinfo.query
//
// WMS 调用该接口，通过 OAID 查询解密后的收件人信息
func Taobaoqimenreceiverinfoquery(clt *core.SDKClient, req *qimen.TaobaoqimenreceiverinfoqueryAPIRequest, session string) (*qimen.TaobaoqimenreceiverinfoqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenreceiverinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
