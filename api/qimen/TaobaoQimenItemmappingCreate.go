package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemmappingCreate 前后端商品映射接口
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
func TaobaoQimenItemmappingCreate(clt *core.SDKClient, req *qimen.TaobaoQimenItemmappingCreateAPIRequest, session string) (*qimen.TaobaoQimenItemmappingCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenItemmappingCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
