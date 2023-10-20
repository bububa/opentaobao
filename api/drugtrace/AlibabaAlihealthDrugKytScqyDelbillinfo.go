package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqydelbillinfo 根据单据号删除单据
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
func Alibabaalihealthdrugkytscqydelbillinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqydelbillinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqydelbillinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqydelbillinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
