package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRightUnbindBack 权限回流-解绑
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
func AlibabaAlihouseNewhomeRightUnbindBack(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRightUnbindBackAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeRightUnbindBackAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeRightUnbindBackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
