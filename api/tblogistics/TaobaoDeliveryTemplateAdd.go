package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaodeliverytemplateadd 新增运费模板
// taobao.delivery.template.add
//
// 增加运费模板的外部接口
func Taobaodeliverytemplateadd(clt *core.SDKClient, req *tblogistics.TaobaodeliverytemplateaddAPIRequest, session string) (*tblogistics.TaobaodeliverytemplateaddAPIResponse, error) {
	var resp tblogistics.TaobaodeliverytemplateaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
