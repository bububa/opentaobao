package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundbatchget 批量获取openmall退款单
// taobao.openmall.refund.batch.get
//
// 批量获取openmall退款单
// 注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
func Taobaoopenmallrefundbatchget(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundbatchgetAPIRequest, session string) (*openmall.TaobaoopenmallrefundbatchgetAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundbatchgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
