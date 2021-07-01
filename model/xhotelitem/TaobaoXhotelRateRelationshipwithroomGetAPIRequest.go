package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateRelationshipwithroomGetAPIRequest
查询rpId API请求
taobao.xhotel.rate.relationshipwithroom.get

某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。 */
type TaobaoXhotelRateRelationshipwithroomGetAPIRequest struct {
	model.Params
	// rpId
	_rpId int64
	// 页数
	_pageNo int64
}

// New
