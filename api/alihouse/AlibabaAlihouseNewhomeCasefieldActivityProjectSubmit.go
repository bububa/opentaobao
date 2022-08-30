package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmit 案场活动楼盘维护
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
func AlibabaAlihouseNewhomeCasefieldActivityProjectSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
