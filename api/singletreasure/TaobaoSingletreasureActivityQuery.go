package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityquery 查询活动列表接口
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
func Taobaosingletreasureactivityquery(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityqueryAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityqueryAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
