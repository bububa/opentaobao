package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderFutureFacescanPut 未来酒店扫脸信息上传
// taobao.xhotel.order.future.facescan.put
//
// 未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
func TaobaoXhotelOrderFutureFacescanPut(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderFutureFacescanPutAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderFutureFacescanPutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
