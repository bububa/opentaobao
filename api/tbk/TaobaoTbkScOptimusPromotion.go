package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscoptimuspromotion 淘宝客-服务商-权益物料精选
// taobao.tbk.sc.optimus.promotion
//
// 服务商使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
func Taobaotbkscoptimuspromotion(clt *core.SDKClient, req *tbk.TaobaotbkscoptimuspromotionAPIRequest, session string) (*tbk.TaobaotbkscoptimuspromotionAPIResponse, error) {
	var resp tbk.TaobaotbkscoptimuspromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
