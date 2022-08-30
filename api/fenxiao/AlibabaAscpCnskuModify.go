package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuModify 供应链中台货品修改接口
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
func AlibabaAscpCnskuModify(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuModifyAPIRequest, session string) (*fenxiao.AlibabaAscpCnskuModifyAPIResponse, error) {
	var resp fenxiao.AlibabaAscpCnskuModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
