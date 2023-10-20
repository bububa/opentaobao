package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalregistrationsyncnew 阿里健康新挂号数据回传
// alibaba.alihealth.medical.registration.syncnew
//
// 阿里健康新挂号记录回传接口
func Alibabaalihealthmedicalregistrationsyncnew(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalregistrationsyncnewAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalregistrationsyncnewAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalregistrationsyncnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
