package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibababaseordersuppliernotify 阿里通信运营商信息回传
// alibaba.base.order.supplier.notify
//
// 接收阿里通信流量运营商信息回传
func Alibababaseordersuppliernotify(clt *core.SDKClient, req *alicom.AlibababaseordersuppliernotifyAPIRequest, session string) (*alicom.AlibababaseordersuppliernotifyAPIResponse, error) {
	var resp alicom.AlibababaseordersuppliernotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
