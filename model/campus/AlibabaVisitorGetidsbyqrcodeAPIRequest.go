package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaVisitorGetidsbyqrcodeAPIRequest
根据访客二维码查访客行程id API请求
alibaba.visitor.getidsbyqrcode

根据支付宝阿里访客小程序的动态二维码查询来访行程id */
type AlibabaVisitorGetidsbyqrcodeAPIRequest struct {
	model.Params
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
	// 来访时间
	_date string
	// 二维码字符串
	_qrCode string
}

// New
