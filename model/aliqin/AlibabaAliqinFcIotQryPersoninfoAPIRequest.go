package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotQryPersoninfoAPIRequest
查询物联卡个人实人认证信息 API请求
alibaba.aliqin.fc.iot.qry.personinfo

查询物联卡个人实人认证信息 */
type AlibabaAliqinFcIotQryPersoninfoAPIRequest struct {
	model.Params
	// 需要查询的iccid
	_iccid string
	// 指定查询某userid
	_userid string
	// 由系统根据业务分配
	_midPatChannel string
}

// New
