package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenChannelinventoryQuery 渠道库存查询接口
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
func TaobaoQimenChannelinventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenChannelinventoryQueryAPIRequest, resp *qimen.TaobaoQimenChannelinventoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
