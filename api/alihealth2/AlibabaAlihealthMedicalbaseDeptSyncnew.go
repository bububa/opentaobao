package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasedeptsyncnew 直连科室上传
// alibaba.alihealth.medicalbase.dept.syncnew
//
// 直连科室上传接口
func Alibabaalihealthmedicalbasedeptsyncnew(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasedeptsyncnewAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasedeptsyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasedeptsyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
