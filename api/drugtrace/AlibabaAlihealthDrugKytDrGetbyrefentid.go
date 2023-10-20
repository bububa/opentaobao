package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrGetbyrefentid 多融通过一个企业唯一标识查询企业详细信息
// alibaba.alihealth.drug.kyt.dr.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugKytDrGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrGetbyrefentidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
