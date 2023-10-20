package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseBusinessActivityDelete 电商券活动删除
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
func AlibabaAlihouseBusinessActivityDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseBusinessActivityDeleteAPIRequest, resp *alihouse.AlibabaAlihouseBusinessActivityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
