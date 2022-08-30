package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuAdd 货品创建
// alibaba.ascp.cnsku.add
//
// 供应链中台货品创建接口
func AlibabaAscpCnskuAdd(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuAddAPIRequest, session string) (*fenxiao.AlibabaAscpCnskuAddAPIResponse, error) {
	var resp fenxiao.AlibabaAscpCnskuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
