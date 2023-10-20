package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytstorebilllist 零售端平台单据查询
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
func Alibabaalihealthdrugkytstorebilllist(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytstorebilllistAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytstorebilllistAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytstorebilllistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
