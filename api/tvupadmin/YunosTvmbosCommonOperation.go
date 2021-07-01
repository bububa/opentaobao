package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvmbosCommonOperation
应用中心通用服务接口
yunos.tvmbos.common.operation

应用中心相关接口的代理 */
func YunosTvmbosCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosTvmbosCommonOperationAPIRequest, session string) (*tvupadmin.YunosTvmbosCommonOperationAPIResponse, error) {
	var resp tvupadmin.YunosTvmbosCommonOperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
