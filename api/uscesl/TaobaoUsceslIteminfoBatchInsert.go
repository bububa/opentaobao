package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslIteminfoBatchInsert 按商家批量写入商品接口
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
func TaobaoUsceslIteminfoBatchInsert(clt *core.SDKClient, req *uscesl.TaobaoUsceslIteminfoBatchInsertAPIRequest, session string) (*uscesl.TaobaoUsceslIteminfoBatchInsertAPIResponse, error) {
	var resp uscesl.TaobaoUsceslIteminfoBatchInsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
