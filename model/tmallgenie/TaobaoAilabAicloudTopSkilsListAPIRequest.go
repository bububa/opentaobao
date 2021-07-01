package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopSkilsListAPIRequest
获取硬件平台设备下挂载的技能列表 API请求
taobao.ailab.aicloud.top.skils.list

提供给在硬件平台接入设备的技能列表 */
type TaobaoAilabAicloudTopSkilsListAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// query(模糊匹配skillName)
	_query string
	// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
	_type string
	// pageNo
	_pageNo int64
	// pageSize
	_pageSize int64
}

// New
