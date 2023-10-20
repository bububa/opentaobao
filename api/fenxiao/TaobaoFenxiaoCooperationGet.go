package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaocooperationget 供应商或分销商获取合作关系信息
// taobao.fenxiao.cooperation.get
//
// 获取供应商的合作关系信息
func Taobaofenxiaocooperationget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaocooperationgetAPIRequest, session string) (*fenxiao.TaobaofenxiaocooperationgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaocooperationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
