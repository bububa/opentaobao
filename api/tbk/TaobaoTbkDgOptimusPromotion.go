package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgoptimuspromotion 淘宝客-推广者-权益物料精选
// taobao.tbk.dg.optimus.promotion
//
// 推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
func Taobaotbkdgoptimuspromotion(clt *core.SDKClient, req *tbk.TaobaotbkdgoptimuspromotionAPIRequest, session string) (*tbk.TaobaotbkdgoptimuspromotionAPIResponse, error) {
	var resp tbk.TaobaotbkdgoptimuspromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
