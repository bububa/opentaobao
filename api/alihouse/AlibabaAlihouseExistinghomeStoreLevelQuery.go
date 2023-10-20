package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreLevelQuery 门店等级评分查询
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
func AlibabaAlihouseExistinghomeStoreLevelQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeStoreLevelQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
