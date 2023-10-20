package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugWxinfoUpload 小程序数据回传
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
func AlibabaAlihealthDrugWxinfoUpload(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
