package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabaascpcnskumodify 供应链中台货品修改接口
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
func Alibabaascpcnskumodify(clt *core.SDKClient, req *fenxiao.AlibabaascpcnskumodifyAPIRequest, session string) (*fenxiao.AlibabaascpcnskumodifyAPIResponse, error) {
	var resp fenxiao.AlibabaascpcnskumodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
