package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

/* AlibabaHappytripTravelSync
差旅申请单同步接口
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中 */
func AlibabaHappytripTravelSync(clt *core.SDKClient, req *happytrip.AlibabaHappytripTravelSyncAPIRequest, session string) (*happytrip.AlibabaHappytripTravelSyncAPIResponse, error) {
	var resp happytrip.AlibabaHappytripTravelSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
