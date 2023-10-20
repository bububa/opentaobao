package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaodeliverytemplateget 获取用户指定运费模板信息
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
func Taobaodeliverytemplateget(clt *core.SDKClient, req *tblogistics.TaobaodeliverytemplategetAPIRequest, session string) (*tblogistics.TaobaodeliverytemplategetAPIResponse, error) {
	var resp tblogistics.TaobaodeliverytemplategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
