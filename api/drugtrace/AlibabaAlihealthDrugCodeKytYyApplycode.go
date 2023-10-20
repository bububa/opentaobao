package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytyyapplycode 医院药品子码申请接口
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
func Alibabaalihealthdrugcodekytyyapplycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytyyapplycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytyyapplycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytyyapplycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
