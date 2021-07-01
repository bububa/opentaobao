package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopLikeFilterAPIRequest
过滤列表歌曲存在于收藏列表的 API请求
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的 */
type TaobaoAilabAicloudTopLikeFilterAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 音频收藏类型, 四种类型：music,children_song,program,story
	_type string
	// 传入的歌曲列表
	_mediaItems []MediaItem
}

// New
