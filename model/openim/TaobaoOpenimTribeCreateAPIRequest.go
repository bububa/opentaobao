package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeCreateAPIRequest
创建群 API请求
taobao.openim.tribe.create

创建一个openim的群 */
type TaobaoOpenimTribeCreateAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群名称
	_tribeName string
	// 群公告
	_notice string
	// 群类型有两种tribe_type = 0  普通群  普通群有管理员角色，对成员加入有权限控制tribe_type = 1  讨论组  讨论组没有管理员，不能解散
	_tribeType int64
	// 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时  该参数为必选tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
	_members []OpenImUser
}

// New
