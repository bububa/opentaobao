package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* AlibabaLsyCrmActivityDataUpdate
私域导购数据回流接口
alibaba.lsy.crm.activity.data.update

私域导购数据回流接口 */
func AlibabaLsyCrmActivityDataUpdate(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityDataUpdateAPIRequest, session string) (*tmallnr.AlibabaLsyCrmActivityDataUpdateAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmActivityDataUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
