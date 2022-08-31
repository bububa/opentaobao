package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseBusinessActivityDelete 电商券活动删除
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
func AlibabaAlihouseBusinessActivityDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseBusinessActivityDeleteAPIRequest, session string) (*alihouse.AlibabaAlihouseBusinessActivityDeleteAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseBusinessActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
