package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTicketProductUploadAPIRequest
【门票API2.0】门票收费项目管理接口 API请求
alitrip.ticket.product.upload

航旅度假新版门票商品（门票收费项目）管理接口：支持门票商品的发布、编辑。如果在ali_product_id下没有发布过门票商品，则系统将判断为新发布商品，否则是编辑已有商品。可以通过辅助查询接口判断是否已在某个ali_product_id下发布过门票商品。
对应新发布商品的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑商品的情况，ali_product_id和out_product_id至少需要填一个，其他参数都是可选，编辑情况支持增量更新（某个参数不传则使用该商品上原有值）。 */
type AlitripTicketProductUploadAPIRequest struct {
	model.Params
	// 新发布门票商品时必填。阿里旅行提供的收费项目编码，与商户收费项目编码进行绑定。注：一个收费项目编码对应了一个淘宝商品
	_aliProductId int64
	// 可选，商户自定义收费项目编码。注：如果商户上传了自定义收费项目编码，则在价格库存同步接口可以使用该自定义编码更新价格库存。
	_outProductId string
	// 可选，商家想要更新映射关系的时候，用于更新商户自定义收费项目编码。
	_updateOutProductId string
	// 商户收费项目名称
	_outProductName string
	// 新发布门票商品时必填。门票商品 库存类型。1、日历库存， 2、非日历库存
	_inventoryType int64
	// inventory_type=2时必填，指定该收费项目在购买时是否需要买家指定入园日期。1、需要，2-不需要
	_needEnterDate int64
	// 新发布门票商品时必填。门票有效期：指定入园日期后 多少天内有效。当为数字时，表示多少天内有效；当为日期时，表示到某日期有效，日期格式：yyyy-MM-dd。发布时不填则默认设置30天内有效
	_expireDate string
	// 新发布门票商品时必填。门票 预定时间限制。1、表示无限制 购买后可立即入园，2、有限制，此时预定时间限制规则必填。
	_reserveLimitType int64
	// 特殊选填，门票 预定时间限制规则。格式：1_18_00_3 或者 1_18_00_3_30，含义：必须提前1天拍下，且在18点00分前支付成功，订单才生效。当为提前0天时（即当日票），最后两个数字才生效，指当日票需要在出票后3小时30分钟后方可入园。
	_reserveLimitRule string
	// 新发布门票商品时必填。门票商品发码方式
	_codeSendingInfo *CodeSendingInfo
	// 可选，商品标题
	_title string
	// 可选，商品主图，最多支持5张。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
	_picUrls []string
	// 可选，商品详情描述，不超过50000个字符。详情描述支持纯文本描述，也支持html格式的详情描述。html格式的详情描述中 图片链接支持外链图片（必须外网可访问， 且格式为png、jpg或jpeg，大小在500k以内）和淘宝图片空间链接。
	_desc string
}

// New
