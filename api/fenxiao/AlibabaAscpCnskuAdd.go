package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabaascpcnskuadd 货品创建
// alibaba.ascp.cnsku.add
//
// 供应链中台货品创建接口
func Alibabaascpcnskuadd(clt *core.SDKClient, req *fenxiao.AlibabaascpcnskuaddAPIRequest, session string) (*fenxiao.AlibabaascpcnskuaddAPIResponse, error) {
	var resp fenxiao.AlibabaascpcnskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
