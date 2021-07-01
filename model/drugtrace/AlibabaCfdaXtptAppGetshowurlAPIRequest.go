package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCfdaXtptAppGetshowurlAPIRequest
协同平台码查询页面url API请求
alibaba.cfda.xtpt.app.getshowurl

协同平台码查询页面url */
type AlibabaCfdaXtptAppGetshowurlAPIRequest struct {
	model.Params
	// 码
	_code string
}

// New
