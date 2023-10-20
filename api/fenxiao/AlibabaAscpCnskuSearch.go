package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabaascpcnskusearch 供应链中台货品搜索接口
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
func Alibabaascpcnskusearch(clt *core.SDKClient, req *fenxiao.AlibabaascpcnskusearchAPIRequest, session string) (*fenxiao.AlibabaascpcnskusearchAPIResponse, error) {
	var resp fenxiao.AlibabaascpcnskusearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
