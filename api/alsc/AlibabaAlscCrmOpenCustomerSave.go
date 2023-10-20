package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopencustomersave 保存和更新顾客
// alibaba.alsc.crm.open.customer.save
//
// 用来保存顾客，如果已经存在的话，则更新顾客
func Alibabaalsccrmopencustomersave(clt *core.SDKClient, req *alsc.AlibabaalsccrmopencustomersaveAPIRequest, session string) (*alsc.AlibabaalsccrmopencustomersaveAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopencustomersaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
