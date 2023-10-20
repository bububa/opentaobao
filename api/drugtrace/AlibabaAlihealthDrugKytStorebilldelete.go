package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytstorebilldelete 零售端单据删除
// alibaba.alihealth.drug.kyt.storebilldelete
//
// 零售端单据删除
func Alibabaalihealthdrugkytstorebilldelete(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytstorebilldeleteAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytstorebilldeleteAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytstorebilldeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
