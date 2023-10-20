package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityitembatchupdate 批量修改商品接口
// taobao.singletreasure.activity.item.batchupdate
//
// 批量修改商品优惠接口
func Taobaosingletreasureactivityitembatchupdate(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityitembatchupdateAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityitembatchupdateAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityitembatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
