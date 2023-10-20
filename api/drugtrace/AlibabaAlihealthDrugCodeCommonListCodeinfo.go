package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeCommonListCodeinfo 通用查询码接口
// alibaba.alihealth.drug.code.common.list.codeinfo
//
// 通用查询码接口
func AlibabaAlihealthDrugCodeCommonListCodeinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeCommonListCodeinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
