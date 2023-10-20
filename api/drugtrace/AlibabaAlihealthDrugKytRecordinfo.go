package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytRecordinfo 快易通健康检查
// alibaba.alihealth.drug.kyt.recordinfo
//
// 快易通健康检查
func AlibabaAlihealthDrugKytRecordinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytRecordinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytRecordinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytRecordinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
