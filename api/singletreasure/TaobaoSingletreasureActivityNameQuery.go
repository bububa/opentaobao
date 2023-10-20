package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivitynamequery 查询官方的活动名称接口
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
func Taobaosingletreasureactivitynamequery(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivitynamequeryAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivitynamequeryAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivitynamequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
