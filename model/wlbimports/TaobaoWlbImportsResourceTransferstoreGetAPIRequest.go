package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsResourceTransferstoreGetAPIRequest
根据指定的资源获取所有中转仓列表 API请求
taobao.wlb.imports.resource.transferstore.get

根据指定的资源获取所有中转仓列表 */
type TaobaoWlbImportsResourceTransferstoreGetAPIRequest struct {
	model.Params
	// 通过taobao.wlb.imports.resource.get接口查询出来的资源ID
	_resourceId int64
	// 卖家发货地址的区域ID，如果不填则为默认发货地址ID
	_fromId int64
	// 商品前台叶子类目ID
	_cids []int64
	// 买家收货地信息
	_toAddress *ReciverAddressDo
}

// New
