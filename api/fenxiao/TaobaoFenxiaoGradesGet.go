package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaogradesget 分销商等级查询
// taobao.fenxiao.grades.get
//
// 根据供应商ID，查询他的分销商等级信息
func Taobaofenxiaogradesget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaogradesgetAPIRequest, session string) (*fenxiao.TaobaofenxiaogradesgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaogradesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
