package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpSuborderEstcontimeModify 向前修改发货时效
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
func AlibabaAscpSuborderEstcontimeModify(clt *core.SDKClient, req *ascp.AlibabaAscpSuborderEstcontimeModifyAPIRequest, session string) (*ascp.AlibabaAscpSuborderEstcontimeModifyAPIResponse, error) {
	var resp ascp.AlibabaAscpSuborderEstcontimeModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
