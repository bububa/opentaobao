package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

/* AlibabaWestcrmActivityListGet
获取活动列表接口
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏 */
func AlibabaWestcrmActivityListGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmActivityListGetAPIRequest, session string) (*westcrm.AlibabaWestcrmActivityListGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmActivityListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
