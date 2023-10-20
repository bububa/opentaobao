package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemQuery 暗拍b2b商品查询
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
func AlibabaIdleTenderBtobItemQuery(clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemQueryAPIRequest, resp *idle.AlibabaIdleTenderBtobItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
