package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaorequisitionsget 合作申请查询
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
func Taobaofenxiaorequisitionsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaorequisitionsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaorequisitionsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaorequisitionsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
