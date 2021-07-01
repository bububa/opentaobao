package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosOsupdateAppversionPublish
发布应用升级
yunos.osupdate.appversion.publish

发布应用升级任务 */
func YunosOsupdateAppversionPublish(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionPublishAPIRequest, session string) (*tvupadmin.YunosOsupdateAppversionPublishAPIResponse, error) {
	var resp tvupadmin.YunosOsupdateAppversionPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
