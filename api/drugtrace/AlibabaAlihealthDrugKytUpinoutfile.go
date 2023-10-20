package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpinoutfile 上传出入库单据(传文件)
// alibaba.alihealth.drug.kyt.upinoutfile
//
// 上传出入库单据(传文件)
func AlibabaAlihealthDrugKytUpinoutfile(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpinoutfileAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpinoutfileAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
