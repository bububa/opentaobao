package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Tmallcityretailfulfillabnormalcenterabnormalstatuschange 同城零售履约异常中心异常单处理结果回调接口
// tmall.cityretail.fulfill.abnormal.center.abnormal.status.change
//
// 同城零售履约异常中心异常单处理结果回调接口
func Tmallcityretailfulfillabnormalcenterabnormalstatuschange(clt *core.SDKClient, req *wdk.TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIRequest, session string) (*wdk.TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIResponse, error) {
	var resp wdk.TmallcityretailfulfillabnormalcenterabnormalstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
