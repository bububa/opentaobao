package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminEpgDesktopOperation 影视桌面运营通用接口
// yunos.tvpubadmin.epg.desktop.operation
//
// 影视桌面运营通用接口
func YunosTvpubadminEpgDesktopOperation(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminEpgDesktopOperationAPIRequest, resp *tvupadmin.YunosTvpubadminEpgDesktopOperationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
