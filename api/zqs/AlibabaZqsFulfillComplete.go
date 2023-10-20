package zqs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/zqs"
)

// AlibabaZqsFulfillComplete 周期购履约完成接口
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
func AlibabaZqsFulfillComplete(clt *core.SDKClient, req *zqs.AlibabaZqsFulfillCompleteAPIRequest, session string) (*zqs.AlibabaZqsFulfillCompleteAPIResponse, error) {
	var resp zqs.AlibabaZqsFulfillCompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
