package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosOsupdateAppversionInfo
获取应用升级详情
yunos.osupdate.appversion.info

获取应用升级详情 */
func YunosOsupdateAppversionInfo(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionInfoAPIRequest, session string) (*tvupadmin.YunosOsupdateAppversionInfoAPIResponse, error) {
	var resp tvupadmin.YunosOsupdateAppversionInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
