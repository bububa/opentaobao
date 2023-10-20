package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice 医疗器械的码查询信息接口
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
func AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldevice(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
