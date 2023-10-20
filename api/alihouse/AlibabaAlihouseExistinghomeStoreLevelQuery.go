package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreLevelQuery 门店等级评分查询
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
func AlibabaAlihouseExistinghomeStoreLevelQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
