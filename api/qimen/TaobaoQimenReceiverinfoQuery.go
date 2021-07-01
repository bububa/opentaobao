package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

/* TaobaoQimenReceiverinfoQuery
OAID 收件人信息解密接口
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息 */
func TaobaoQimenReceiverinfoQuery(clt *core.SDKClient, req *qimen.TaobaoQimenReceiverinfoQueryAPIRequest, session string) (*qimen.TaobaoQimenReceiverinfoQueryAPIResponse, error) {
	var resp qimen.TaobaoQimenReceiverinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
