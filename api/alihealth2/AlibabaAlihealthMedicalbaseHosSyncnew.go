package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasehossyncnew 直连医院上传接口
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
func Alibabaalihealthmedicalbasehossyncnew(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasehossyncnewAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasehossyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasehossyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
