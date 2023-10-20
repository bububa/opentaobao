package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytupdatebillcode 零售修改出入库单追溯码
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
func Alibabaalihealthdrugkytupdatebillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytupdatebillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytupdatebillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytupdatebillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
