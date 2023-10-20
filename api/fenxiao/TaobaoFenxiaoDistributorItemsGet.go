package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodistributoritemsget 查询商品下载记录
// taobao.fenxiao.distributor.items.get
//
// 供应商查询分销商商品下载记录。
func Taobaofenxiaodistributoritemsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodistributoritemsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaodistributoritemsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodistributoritemsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
