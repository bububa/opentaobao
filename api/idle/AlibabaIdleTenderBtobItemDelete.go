package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemDelete 暗拍b2b商品下架/删除
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
func AlibabaIdleTenderBtobItemDelete(clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemDeleteAPIRequest, resp *idle.AlibabaIdleTenderBtobItemDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
