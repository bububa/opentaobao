package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminepgdesktopoperation 影视桌面运营通用接口
// yunos.tvpubadmin.epg.desktop.operation
//
// 影视桌面运营通用接口
func Yunostvpubadminepgdesktopoperation(clt *core.SDKClient, req *tvupadmin.YunostvpubadminepgdesktopoperationAPIRequest, session string) (*tvupadmin.YunostvpubadminepgdesktopoperationAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminepgdesktopoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
