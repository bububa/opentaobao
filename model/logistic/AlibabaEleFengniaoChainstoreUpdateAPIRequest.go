package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreUpdateAPIRequest
修改门店信息接口 API请求
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名 */
type AlibabaEleFengniaoChainstoreUpdateAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// New
