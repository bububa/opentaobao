package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* TaobaoXhotelOrderFutureInfoPut
订单信息上传更新
taobao.xhotel.order.future.info.put

商家调用推送信息给飞猪平台。 支持如下操作类型：21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统） */
func TaobaoXhotelOrderFutureInfoPut(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderFutureInfoPutAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderFutureInfoPutAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelOrderFutureInfoPutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
