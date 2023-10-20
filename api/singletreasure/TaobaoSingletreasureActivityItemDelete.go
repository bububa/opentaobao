package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityitemdelete 删除单品优惠接口
// taobao.singletreasure.activity.item.delete
//
// 删除单品优惠接口
func Taobaosingletreasureactivityitemdelete(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityitemdeleteAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityitemdeleteAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
