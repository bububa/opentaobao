package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivitycreate 活动创建接口
// taobao.singletreasure.activity.create
//
// 创建优惠活动
func Taobaosingletreasureactivitycreate(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivitycreateAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivitycreateAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
