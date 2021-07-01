package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopLikeAddAPIRequest
增加收藏 API请求
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏 */
type TaobaoAilabAicloudTopLikeAddAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离
	_schema string
	// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
	_type string
	// 来源
	_source string
	// 收藏的资源的ID
	_itemId string
	// 内容，必须要是一个json格式：{"song":"走过1999","singer":"张学友","album":"走过1999"}
	_content string
}

// New
