package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// Taobaosingletreasureactivityitemquery 查询活动下的优惠信息
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
func Taobaosingletreasureactivityitemquery(clt *core.SDKClient, req *singletreasure.TaobaosingletreasureactivityitemqueryAPIRequest, session string) (*singletreasure.TaobaosingletreasureactivityitemqueryAPIResponse, error) {
	var resp singletreasure.TaobaosingletreasureactivityitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
