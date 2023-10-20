package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbhouseAdd 民宿门店信息添加
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
func TaobaoXhotelBnbhouseAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbhouseAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbhouseAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
