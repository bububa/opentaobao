package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMemberinfoUpdateAPIRequest
编辑会员资料 API请求
taobao.crm.memberinfo.update

编辑会员的基本资料，接口返回会员信息修改是否成功 */
type TaobaoCrmMemberinfoUpdateAPIRequest struct {
	model.Params
	// 买家昵称
	_buyerNick string
	// 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).
	_status string
	// 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效
	_grade int64
	// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
	_province string
	// 城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线
	_city string
	// 交易笔数
	_tradeCount int64
	// 交易金额，单位：分
	_tradeAmount int64
	// 交易关闭次数
	_closeTradeCount int64
	// 交易关闭金额，单位：分
	_closeTradeAmount int64
	// 分组的id集合字符串
	_groupIds string
	// 宝贝件数
	_itemNum int64
}

// New
