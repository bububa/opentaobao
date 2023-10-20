package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytupstorebillfile 上传零售出入库单(上传文件)
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
func Alibabaalihealthdrugkytupstorebillfile(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytupstorebillfileAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytupstorebillfileAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytupstorebillfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
