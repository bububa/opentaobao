package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice 医疗器械的码查询信息接口
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
func AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
