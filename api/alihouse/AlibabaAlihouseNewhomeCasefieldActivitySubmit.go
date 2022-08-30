package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmit 案场活动维护
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
func AlibabaAlihouseNewhomeCasefieldActivitySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
