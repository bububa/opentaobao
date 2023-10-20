package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytupinoutfile 上传出入库单据(传文件)
// alibaba.alihealth.drug.kyt.upinoutfile
//
// 上传出入库单据(传文件)
func Alibabaalihealthdrugkytupinoutfile(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytupinoutfileAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytupinoutfileAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytupinoutfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
