package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabaascpcnskuupdate 供应链中台货品修改接口
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
func Alibabaascpcnskuupdate(clt *core.SDKClient, req *fenxiao.AlibabaascpcnskuupdateAPIRequest, session string) (*fenxiao.AlibabaascpcnskuupdateAPIResponse, error) {
	var resp fenxiao.AlibabaascpcnskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
