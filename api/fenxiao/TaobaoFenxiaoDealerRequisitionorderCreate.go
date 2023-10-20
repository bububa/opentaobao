package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodealerrequisitionordercreate 创建经销采购申请
// taobao.fenxiao.dealer.requisitionorder.create
//
// 创建经销采购申请
func Taobaofenxiaodealerrequisitionordercreate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodealerrequisitionordercreateAPIRequest, session string) (*fenxiao.TaobaofenxiaodealerrequisitionordercreateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodealerrequisitionordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
