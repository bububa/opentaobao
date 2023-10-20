package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelorderfuturefacescanput 未来酒店扫脸信息上传
// taobao.xhotel.order.future.facescan.put
//
// 未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
func Taobaoxhotelorderfuturefacescanput(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelorderfuturefacescanputAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelorderfuturefacescanputAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelorderfuturefacescanputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
