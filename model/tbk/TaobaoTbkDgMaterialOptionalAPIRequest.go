package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgMaterialOptionalAPIRequest
淘宝客-推广者-物料搜索 API请求
taobao.tbk.dg.material.optional

通用物料搜索API（导购） */
type TaobaoTbkDgMaterialOptionalAPIRequest struct {
	model.Params
	// 商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	_startDsr int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 第几页，默认：１
	_pageNo int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 商品筛选-淘客佣金比率上限。如：1234表示12.34%
	_endTkRate int64
	// 商品筛选-淘客佣金比率下限。如：1234表示12.34%
	_startTkRate int64
	// 商品筛选-折扣价范围上限。单位：元
	_endPrice int64
	// 商品筛选-折扣价范围下限。单位：元
	_startPrice int64
	// 商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	_isOverseas bool
	// 商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	_isTmall bool
	// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）
	_sort string
	// 商品筛选-所在地
	_itemloc string
	// 商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	_cat string
	// 商品筛选-查询词
	_q string
	// 不传时默认物料id=2836；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004
	_materialId int64
	// 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	_hasCoupon bool
	// ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_ip string
	// mm_xxx_xxx_12345678三段式的最后一段数字
	_adzoneId int64
	// 商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	_needFreeShipment bool
	// 商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	_needPrepay bool
	// 商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includePayRate30 bool
	// 商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	_includeGoodRate bool
	// 商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	_includeRfdRate bool
	// 商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	_npxLevel int64
	// 商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
	_endKaTkRate int64
	// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	_startKaTkRate int64
	// 智能匹配-设备号加密类型：MD5
	_deviceEncrypt string
	// 智能匹配-设备号加密后的值（MD5加密需32位小写）
	_deviceValue string
	// 智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	_deviceType string
	// 锁佣结束时间
	_lockRateEndTime int64
	// 锁佣开始时间
	_lockRateStartTime int64
	// 本地化业务入参-LBS信息-经度
	_longitude string
	// 本地化业务入参-LBS信息-纬度
	_latitude string
	// 本地化业务入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）
	_cityCode string
	// 商家id，仅支持饿了么卡券商家ID，支持批量请求1-100以内，多个商家ID使用英文逗号分隔
	_sellerIds string
	// 会员运营ID
	_specialId string
	// 渠道关系ID，仅适用于渠道推广场景
	_relationId string
	// 本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值
	_pageResultKey string
	// 人群ID，仅适用于物料评估场景material_id=41377
	_ucrowdId int64
	// 物料评估-商品列表
	_ucrowdRankItems []Ucrowdrankitems
	// 是否获取前N件佣金信息	0否，1是，其他值否
	_getTopnRate int64
}

// New
