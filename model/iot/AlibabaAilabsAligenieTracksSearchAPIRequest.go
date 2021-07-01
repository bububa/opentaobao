package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieTracksSearchAPIRequest
查询音频 API请求
alibaba.ailabs.aligenie.tracks.search

搜索类目下的音频信息 */
type AlibabaAilabsAligenieTracksSearchAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 一级类目，如：儿童、新闻、商业财经有声书等
	_param1 string
	// 二级类目，如儿童下有：儿歌、童谣、国学等
	_param2 string
	// 搜索的单个音频名称
	_param3 string
	// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
	_param4 int64
	// 每页数量（不超过50）
	_param5 int64
}

// New
