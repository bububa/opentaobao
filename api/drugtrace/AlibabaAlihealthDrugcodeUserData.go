package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodeuserdata 西安杨森同步用户行为接口
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
func Alibabaalihealthdrugcodeuserdata(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodeuserdataAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodeuserdataAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodeuserdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
