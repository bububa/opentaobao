package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabaascpsuborderestcontimemodify 向前修改发货时效
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
func Alibabaascpsuborderestcontimemodify(clt *core.SDKClient, req *ascp.AlibabaascpsuborderestcontimemodifyAPIRequest, session string) (*ascp.AlibabaascpsuborderestcontimemodifyAPIResponse, error) {
	var resp ascp.AlibabaascpsuborderestcontimemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
