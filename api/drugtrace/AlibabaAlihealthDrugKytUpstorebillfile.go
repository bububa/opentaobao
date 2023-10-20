package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpstorebillfile 上传零售出入库单(上传文件)
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
func AlibabaAlihealthDrugKytUpstorebillfile(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpstorebillfileAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
