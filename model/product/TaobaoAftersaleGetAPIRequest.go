package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAftersaleGetAPIRequest
查询用户售后服务模板 API请求
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id */
type TaobaoAftersaleGetAPIRequest struct {
	model.Params
}

// New
