package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvmboscommonoperation 应用中心通用服务接口
// yunos.tvmbos.common.operation
//
// 应用中心相关接口的代理
func Yunostvmboscommonoperation(clt *core.SDKClient, req *tvupadmin.YunostvmboscommonoperationAPIRequest, session string) (*tvupadmin.YunostvmboscommonoperationAPIResponse, error) {
	var resp tvupadmin.YunostvmboscommonoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
