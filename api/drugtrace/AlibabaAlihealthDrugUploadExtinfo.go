package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdruguploadextinfo 中药饮片及器械对接
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
func Alibabaalihealthdruguploadextinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdruguploadextinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdruguploadextinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdruguploadextinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
