package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgOptimusMaterialAPIRequest
淘宝客-推广者-物料精选 API请求
taobao.tbk.dg.optimus.material

支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin） */
type TaobaoTbkDgOptimusMaterialAPIRequest struct {
	model.Params
	// 页大小，默认20，1~100
	_pageSize int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 第几页，默认：1
	_pageNo int64
	// 官方的物料Id(详细物料id见：https://market.m.taobao.com/app/qn/toutiao-new/index-pc.html#/detail/10628875?_k=gpov9a)
	_materialId int64
	// 智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	_deviceValue string
	// 智能匹配-设备号加密类型：MD5，类型为OAID时不传
	_deviceEncrypt string
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_deviceType string
	// 内容专用-内容详情ID
	_contentId int64
	// 内容专用-内容渠道信息
	_contentSource string
	// 商品ID，用于相似商品推荐
	_itemId int64
	// 选品库投放id
	_favoritesId string
}

// New
