package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerBillUploadAPIRequest
上传入出库单api API请求
alibaba.alihealth.tracecodeseller.bill.upload

上传入出库单api */
type AlibabaAlihealthTracecodesellerBillUploadAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 单据编号
	_billCode string
	// 出入库类型
	_type string
	// 出入库时间 精确到 年 月 日 时 分 秒
	_time string
	// 自己仓库id
	_warehouseId int64
	// 渠道商id
	_entMerchantId int64
	// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
	_codeInfo string
}

// New
