package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeEcodeUpdate 新房货变更E码
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
func AlibabaAlihouseNewhomeEcodeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeEcodeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
