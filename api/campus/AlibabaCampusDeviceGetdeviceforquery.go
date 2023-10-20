package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdevicegetdeviceforquery 下发设备的分页接口(无需AOP控制)
// alibaba.campus.device.getdeviceforquery
//
// 下发设备的分页接口(发布在TOP上，connect调用，无需AOP控制)
func Alibabacampusdevicegetdeviceforquery(clt *core.SDKClient, req *campus.AlibabacampusdevicegetdeviceforqueryAPIRequest, session string) (*campus.AlibabacampusdevicegetdeviceforqueryAPIResponse, error) {
	var resp campus.AlibabacampusdevicegetdeviceforqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
