package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqylistcodefullinfodtomedicaldevice 医疗器械的码查询信息接口
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
func Alibabaalihealthdrugkytscqylistcodefullinfodtomedicaldevice(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqylistcodefullinfodtomedicaldeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
