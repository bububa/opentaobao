package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradebatchget 批量获取openmall订单
// taobao.openmall.trade.batch.get
//
// 批量获取openmall订单
// 注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
func Taobaoopenmalltradebatchget(clt *core.SDKClient, req *openmall.TaobaoopenmalltradebatchgetAPIRequest, session string) (*openmall.TaobaoopenmalltradebatchgetAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradebatchgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
