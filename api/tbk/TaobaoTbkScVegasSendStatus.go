package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScVegasSendStatus 淘宝客-服务商-红包领取状态查询
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
func TaobaoTbkScVegasSendStatus(clt *core.SDKClient, req *tbk.TaobaoTbkScVegasSendStatusAPIRequest, session string) (*tbk.TaobaoTbkScVegasSendStatusAPIResponse, error) {
	var resp tbk.TaobaoTbkScVegasSendStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
