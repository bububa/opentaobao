package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityupdate 修改活动接口
// taobao.singletreasure.activity.update
//
// 修改活动接口
func Taobaosingletreasureactivityupdate(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityupdateAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityupdateAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
