package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Alibabainteractsensorsocial 社交组件
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
func Alibabainteractsensorsocial(clt *core.SDKClient, req *mtopopen.AlibabainteractsensorsocialAPIRequest, session string) (*mtopopen.AlibabainteractsensorsocialAPIResponse, error) {
	var resp mtopopen.AlibabainteractsensorsocialAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
