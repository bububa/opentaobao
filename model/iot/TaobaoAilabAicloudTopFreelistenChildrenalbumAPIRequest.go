package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest
儿童音频列表 API请求
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表 */
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
	_param1 string
	// 页数
	_param2 int64
	// 每页条目数
	_param3 int64
}

// New
