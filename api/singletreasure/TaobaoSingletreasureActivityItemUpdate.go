package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityitemupdate 更新单品优惠接口
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
func Taobaosingletreasureactivityitemupdate(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityitemupdateAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityitemupdateAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityitemupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
