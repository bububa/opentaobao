package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRightBindBack 权限回流
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
func AlibabaAlihouseNewhomeRightBindBack(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRightBindBackAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeRightBindBackAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeRightBindBackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
