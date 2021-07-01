package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemMapDeleteAPIRequest
失效指定用户的商品与后端商品的映射关系 API请求
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系 */
type TaobaoScitemMapDeleteAPIRequest struct {
	model.Params
	// 后台商品ID
	_scItemId int64
	// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
	_userNick string
}

// New
