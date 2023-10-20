package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampuscorecompanycampusgetcombycamid 根据园区ID获取运营公司信息
// alibaba.campus.core.companycampus.getcombycamid
//
// 根据园区ID获取运营公司信息
func Alibabacampuscorecompanycampusgetcombycamid(clt *core.SDKClient, req *campus.AlibabacampuscorecompanycampusgetcombycamidAPIRequest, session string) (*campus.AlibabacampuscorecompanycampusgetcombycamidAPIResponse, error) {
	var resp campus.AlibabacampuscorecompanycampusgetcombycamidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
