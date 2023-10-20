package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaodeliverytemplatedelete 删除运费模板
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
func Taobaodeliverytemplatedelete(clt *core.SDKClient, req *tblogistics.TaobaodeliverytemplatedeleteAPIRequest, session string) (*tblogistics.TaobaodeliverytemplatedeleteAPIResponse, error) {
	var resp tblogistics.TaobaodeliverytemplatedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
