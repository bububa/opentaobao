package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBnbroomtypeAddAPIRequest
民宿新增房源 API请求
taobao.xhotel.bnbroomtype.add

添加民宿房源 */
type TaobaoXhotelBnbroomtypeAddAPIRequest struct {
	model.Params
	// 销售渠道,默认taobao
	_vendor string
	// 房型id, 这是卖家自己系统中的ID
	_outerId string
	// 外部门店id
	_outHid string
	// 房型名
	_name string
	// 房型英文名
	_nameE string
	// 房型类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_productType int64
	// 有无资质执照 0 没有 1有
	_hasLicense int64
	// 单间面积，单位平方米
	_houseSize int64
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 房型介绍
	_introduction string
	// 亮点描述
	_brightspot string
	// 位置描述
	_localInfo string
	// 周边描述
	_surroundInfo string
	// 品牌名称
	_brand string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 装修时间，格式为2015-01-01装修时间
	_decorateTime string
	// 装修等级 1 精装；2普通；3简装
	_decorateLevel int64
	// 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_decorateStyle int64
	// 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_scenicFeature int64
	// 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
	_rentType int64
	// 单间面积，单位平方米
	_rentSize int64
	// 是否与房东同住 0 不同住 1同住
	_hasLandlord int64
	// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 真实联系方式
	_realTel string
	// 状态 0：在线 -1：不在线 -2:停售
	_status *model.File
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 是否支持IM聊天 0不支持 1支持
	_supportIm int64
	// 是否开启闪订 0不开启 1开启
	_quickOrder int64
	// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
	_bedInfo string
	// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
	_houseModel string
	// 窗型-1.有窗；2.无窗；3.部分有窗
	_windowType int64
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics []BnbPictureDto
	// 房型外部标签 标签信息，逗号(,)分隔
	_outerTags string
	// 是否使用实拍图片 0不使用 1使用
	_isUseShootImage int64
	// 视频地址
	_videoUrl string
	// 民宿房源位置信息
	_location *BnbLocationDto
	// 最大入住人数 1-50
	_maxOccupancy int64
	// 民宿入住要求&附加信息
	_bnbBookingTime *BnbBookingTimeDto
	// 入住须知
	_checkInNotes string
	// 0：不限制，1：只限男性，2：只限女性'
	_guestGender int64
	// 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_guestAge int64
	// 是否可接待外宾 0：否 1：是
	_receiveForeigners int64
	// “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_cleaningFrequency int64
	// 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_activitiesAllowed string
	// 0-n；若不可加床，值为0
	_extraBedsNum int64
	// 押金类型0.线下；1.线上
	_depositType int64
	// 是否信用免押金0：否 1：是
	_supportcredit int64
	// 押金金额
	_depositAmount int64
	// 加人收费信息
	_charge *BnbChargeDto
	// 清洁费是否收取 0：否 1：是
	_cleaningCharge int64
	// 清洁费类型 0.线下；1.线上
	_cleaningType int64
	// 清洁费金额；整数[1，9999999]
	_extraCleaningCharge int64
	// 发票，0：卖家提供发票，1：房东提供发票
	_invoice int64
	// 可提供发票类型，1.专票 2.纸质普票 3.电子普票
	_invoiceType int64
	// 是否有前台 0没有 1有
	_hasFrontDesk int64
	// 如果要变更商品房型编码请使用该字段。
	_newOuterId string
	// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}，见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_service string
}

// New
