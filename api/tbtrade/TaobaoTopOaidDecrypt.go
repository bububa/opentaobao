package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotopoaiddecrypt OAID解密
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
func Taobaotopoaiddecrypt(clt *core.SDKClient, req *tbtrade.TaobaotopoaiddecryptAPIRequest, session string) (*tbtrade.TaobaotopoaiddecryptAPIResponse, error) {
	var resp tbtrade.TaobaotopoaiddecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
