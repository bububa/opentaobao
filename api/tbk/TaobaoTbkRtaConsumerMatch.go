package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkrtaconsumermatch 淘宝客-推广者-定向活动目标发布
// taobao.tbk.rta.consumer.match
//
// 淘客计划向用户推送某个定向活动时，调用该接口判断用户是否符合活动目标（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
func Taobaotbkrtaconsumermatch(clt *core.SDKClient, req *tbk.TaobaotbkrtaconsumermatchAPIRequest, session string) (*tbk.TaobaotbkrtaconsumermatchAPIResponse, error) {
	var resp tbk.TaobaotbkrtaconsumermatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
