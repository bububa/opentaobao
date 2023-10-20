package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityitembatchadd 批量添加商品接口
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
func Taobaosingletreasureactivityitembatchadd(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityitembatchaddAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityitembatchaddAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityitembatchaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
