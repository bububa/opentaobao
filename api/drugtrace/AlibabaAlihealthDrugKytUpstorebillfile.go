package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpstorebillfile 上传零售出入库单(上传文件)
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
func AlibabaAlihealthDrugKytUpstorebillfile(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
