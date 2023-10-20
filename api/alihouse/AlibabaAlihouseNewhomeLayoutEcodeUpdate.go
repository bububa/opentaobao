package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdate 新房户型变更E码
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
func AlibabaAlihouseNewhomeLayoutEcodeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
