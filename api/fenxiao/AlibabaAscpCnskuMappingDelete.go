package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Alibabaascpcnskumappingdelete 货品关系解绑
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
func Alibabaascpcnskumappingdelete(clt *core.SDKClient, req *fenxiao.AlibabaascpcnskumappingdeleteAPIRequest, session string) (*fenxiao.AlibabaascpcnskumappingdeleteAPIResponse, error) {
	var resp fenxiao.AlibabaascpcnskumappingdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
