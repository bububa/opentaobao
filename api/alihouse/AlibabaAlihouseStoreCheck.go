package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseStoreCheck 门店对账查询工具
// alibaba.alihouse.store.check
//
// 门店对账查询工具
func AlibabaAlihouseStoreCheck(clt *core.SDKClient, req *alihouse.AlibabaAlihouseStoreCheckAPIRequest, session string) (*alihouse.AlibabaAlihouseStoreCheckAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseStoreCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
