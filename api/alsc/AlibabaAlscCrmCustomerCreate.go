package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomercreate 创建顾客
// alibaba.alsc.crm.customer.create
//
// 开放本地生活创建顾客功能
func Alibabaalsccrmcustomercreate(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomercreateAPIRequest, session string) (*alsc.AlibabaalsccrmcustomercreateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
