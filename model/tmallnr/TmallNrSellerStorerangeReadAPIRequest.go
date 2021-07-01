package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrSellerStorerangeReadAPIRequest
门店服务范围读取 API请求
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围 */
type TmallNrSellerStorerangeReadAPIRequest struct {
	model.Params
	// 给ISV用sellerid，ISV可能对接其他seller，有可能和登录不是同一个sellerid
	_sellerId int64
	// 业务身份，此api非必填
	_bizIdentity string
	// 门店id
	_storeIds []int64
}

// New
