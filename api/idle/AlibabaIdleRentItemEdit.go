package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemEdit 租赁商品编辑
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
func AlibabaIdleRentItemEdit(clt *core.SDKClient, req *idle.AlibabaIdleRentItemEditAPIRequest, session string) (*idle.AlibabaIdleRentItemEditAPIResponse, error) {
	var resp idle.AlibabaIdleRentItemEditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
