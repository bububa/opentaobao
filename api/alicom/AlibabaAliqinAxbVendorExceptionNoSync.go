package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinaxbvendorexceptionnosync 中心化供应商异常号码状态同步接口
// alibaba.aliqin.axb.vendor.exception.no.sync
//
// 用于中心化供应商同步异常号码
func Alibabaaliqinaxbvendorexceptionnosync(clt *core.SDKClient, req *alicom.AlibabaaliqinaxbvendorexceptionnosyncAPIRequest, session string) (*alicom.AlibabaaliqinaxbvendorexceptionnosyncAPIResponse, error) {
	var resp alicom.AlibabaaliqinaxbvendorexceptionnosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
