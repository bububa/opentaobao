package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBnbhouseAddAPIRequest
民宿门店信息添加 API请求
taobao.xhotel.bnbhouse.add

添加和更新民宿门店的信息 */
type TaobaoXhotelBnbhouseAddAPIRequest struct {
	model.Params
	// 外部房东id
	_outOwnerId string
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 供应商渠道门店ID
	_outerId string
	// 门店名称
	_name string
	// 门店英文名称
	_nameE string
	// 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
	_attributes int64
	// 门店类型，详见“房源类型list
	_productType int64
	// 有无资质执照 0 无资质 1有资质
	_hasLicense int64
	// 面积大小
	_houseSize int64
	// 楼层
	_floors string
	// 门店简介
	_description string
	// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
	_facilities string
	// 品牌名称
	_brand string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 装修等级 1 精装 2普通 3简装
	_decorateLevel int64
	// 装修时间，格式为2015-01-01
	_decorateTime string
	// 装修风格，详见装修风格枚举表
	_decorateStyle int64
	// 风景类型，详见风景类型枚举表
	_scenicFeature int64
	// 房型状态。0:正常，-1:删除，-2:停售
	_status *model.File
	// 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 真实联系方式
	_realTel string
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 民宿门店添加
	_pictures []BnbPictureDto
	// 门店标签 标签信息，逗号(,)分隔
	_tags string
	// 是否使用实拍图片 0不使用 1使用
	_isUseShootImage int64
	// 视频地址
	_videoUrl string
	// 是否有前台 0没有 1有
	_hasFrontDesk int64
	// 位置信息
	_location *BnbLocationDto
	// 入住要求&附加信息
	_bnbBookingTime *BnbBookingTimeDto
	// 入住须知
	_checkInNotes string
	// 可接待客人性别 0：不限制，1：只限男性，2：只限女性'
	_guestGender int64
	// 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
	_guestAge int64
	// 是否可接待外宾 0不接待 1接待
	_receiveForeigners int64
	// 详见“允许活动”list 12,32,33,
	_activitiesAllowed string
	// 可加床数
	_extraBedsNum int64
	// 加人收费信息
	_charge *BnbChargeDto
}

// New
