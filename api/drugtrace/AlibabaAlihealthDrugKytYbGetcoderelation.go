package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytybgetcoderelation 医保-查询码的所有子码
// alibaba.alihealth.drug.kyt.yb.getcoderelation
//
// 应用于药店或医院入库环节，通过扫码获取下级码进行入库；
// 通过码查询所有子码以及包装比例
func Alibabaalihealthdrugkytybgetcoderelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytybgetcoderelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytybgetcoderelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytybgetcoderelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
