package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivitydelete 删除活动接口
// taobao.singletreasure.activity.delete
//
// 删除优惠活动
func Taobaosingletreasureactivitydelete(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivitydeleteAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivitydeleteAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
