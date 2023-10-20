package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvscreenadmincommonoperation 一体机桌面通用接口
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
func Yunostvscreenadmincommonoperation(clt *core.SDKClient, req *tvupadmin.YunostvscreenadmincommonoperationAPIRequest, session string) (*tvupadmin.YunostvscreenadmincommonoperationAPIResponse, error) {
	var resp tvupadmin.YunostvscreenadmincommonoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
