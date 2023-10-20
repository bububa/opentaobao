package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizstoreinsert 新增电子价签商家门店接口
// taobao.uscesl.biz.store.insert
//
// 新增电子价签商家门店接口
func Taobaousceslbizstoreinsert(clt *core.SDKClient, req *uscesl.TaobaousceslbizstoreinsertAPIRequest, session string) (*uscesl.TaobaousceslbizstoreinsertAPIResponse, error) {
	var resp uscesl.TaobaousceslbizstoreinsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
