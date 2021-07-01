package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieAlbumsGetAPIRequest
专辑详情 API请求
alibaba.ailabs.aligenie.albums.get

给予厂商查询专辑下的音频详情 */
type AlibabaAilabsAligenieAlbumsGetAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 专辑 id
	_param1 int64
	// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
	_param2 int64
	// 每页数量（不超过50）
	_param3 int64
}

// New
