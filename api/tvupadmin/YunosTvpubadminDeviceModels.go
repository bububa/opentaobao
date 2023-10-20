package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindevicemodels 获取品牌下设备列表
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
func Yunostvpubadmindevicemodels(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindevicemodelsAPIRequest, session string) (*tvupadmin.YunostvpubadmindevicemodelsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindevicemodelsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
