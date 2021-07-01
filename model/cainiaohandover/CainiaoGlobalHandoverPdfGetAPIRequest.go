package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverPdfGetAPIRequest
获取面单PDF文件数据 API请求
cainiao.global.handover.pdf.get

返回指定大包面单的PDF文件数据 */
type CainiaoGlobalHandoverPdfGetAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 大包编号id
	_handoverContentId int64
	// 打印数据类型，1：面单、4：发货标签、512：交接清单
	_type int64
}

// New
