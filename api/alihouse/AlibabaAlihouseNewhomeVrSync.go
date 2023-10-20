package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomevrsync VR关系数据同步
// alibaba.alihouse.newhome.vr.sync
//
// 对接易居VR关系数据迁移
func Alibabaalihousenewhomevrsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomevrsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomevrsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomevrsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
