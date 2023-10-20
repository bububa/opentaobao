package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytupdatebillinfo 零售端平台单据更新
// alibaba.alihealth.drug.kyt.updatebillinfo
//
// 零售端平台单据更新
func Alibabaalihealthdrugkytupdatebillinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytupdatebillinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytupdatebillinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytupdatebillinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
