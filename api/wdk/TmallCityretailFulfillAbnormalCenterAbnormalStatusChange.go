package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TmallCityretailFulfillAbnormalCenterAbnormalStatusChange 同城零售履约异常中心异常单处理结果回调接口
// tmall.cityretail.fulfill.abnormal.center.abnormal.status.change
//
// 同城零售履约异常中心异常单处理结果回调接口
func TmallCityretailFulfillAbnormalCenterAbnormalStatusChange(ctx context.Context, clt *core.SDKClient, req *wdk.TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest, resp *wdk.TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
