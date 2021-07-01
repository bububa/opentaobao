package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkPosItemprodFindterminalAPIRequest
终端配置支持 API请求
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回 */
type TaobaoLarkPosItemprodFindterminalAPIRequest struct {
	model.Params
	// 终端id
	_deviceId string
	// 终端类型
	_deviceType string
	// 912874323429834
	_createUser string
	// 租户编码
	_leaseCode string
	// 影城id
	_cinemaId string
	// 影城名称
	_cinemaName string
}

// New
