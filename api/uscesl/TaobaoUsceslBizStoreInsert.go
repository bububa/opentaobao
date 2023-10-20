package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizStoreInsert 新增电子价签商家门店接口
// taobao.uscesl.biz.store.insert
//
// 新增电子价签商家门店接口
func TaobaoUsceslBizStoreInsert(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizStoreInsertAPIRequest, resp *uscesl.TaobaoUsceslBizStoreInsertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
