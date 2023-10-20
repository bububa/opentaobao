package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadinsign 仓库批量扫码回传接口
// alibaba.alihealth.drug.kyt.uploadinsign
//
// 连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
func AlibabaAlihealthDrugKytUploadinsign(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadinsignAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUploadinsignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
