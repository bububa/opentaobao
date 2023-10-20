package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemmapadd 创建IC商品与后端商品的映射关系
// taobao.scitem.map.add
//
// 创建IC商品或分销商品与后端商品的映射关系
func Taobaoscitemmapadd(clt *core.SDKClient, req *fenxiao.TaobaoscitemmapaddAPIRequest, session string) (*fenxiao.TaobaoscitemmapaddAPIResponse, error) {
	var resp fenxiao.TaobaoscitemmapaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
