package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaooctradetracealertsget 异常订单信息获取
// taobao.oc.tradetrace.alerts.get
//
// 提供订单预警模块的数据查询接口
func Taobaooctradetracealertsget(clt *core.SDKClient, req *jst.TaobaooctradetracealertsgetAPIRequest, session string) (*jst.TaobaooctradetracealertsgetAPIResponse, error) {
	var resp jst.TaobaooctradetracealertsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
