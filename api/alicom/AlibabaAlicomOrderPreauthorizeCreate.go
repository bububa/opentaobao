package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomorderpreauthorizecreate 业务办理结果
// alibaba.alicom.order.preauthorize.create
//
// 授授权:签约结果通知
func Alibabaalicomorderpreauthorizecreate(clt *core.SDKClient, req *alicom.AlibabaalicomorderpreauthorizecreateAPIRequest, session string) (*alicom.AlibabaalicomorderpreauthorizecreateAPIResponse, error) {
	var resp alicom.AlibabaalicomorderpreauthorizecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
