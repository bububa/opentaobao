package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecAddAPIRequest
添加产品规格 API请求
tmall.product.spec.add

增加产品规格 */
type TmallProductSpecAddAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本
	_certifiedPicStr string
	// 产品规格吊牌价，以分为单位，无默认值，上限999999999
	_labelPrice int64
	// 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值
	_certifiedTxtStr string
	// 产品的规格属性
	_specProps string
	// 规格属性别名,只允许传颜色别名
	_specPropsAlias string
	// 用户自定义销售属性，结构：pid1:value1;pid2:value2。在
	_customerSpecProps string
	// 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1
	_changeProp string
	// 产品图片
	_image *model.File
	// 产品二维码
	_barcode string
	// 产品货号
	_productCode string
	// 产品上市时间
	_marketTime string
}

// New
