package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareAddowAPIRequest
【国际机票自有政策】单条单程添加 API请求
taobao.alitrip.it.fare.addow

自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel */
type TaobaoAlitripItFareAddowAPIRequest struct {
	model.Params
	// 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
	_outFileCode string
	// 文件编号
	_fileCode string
	// （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
	_productType string
	// （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
	_stockMode string
	// 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
	_ticketingAirline string
	// 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
	_saleAirline string
	// 城市/机场选项，默认为城市1、可以填写：“机场",“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
	_addressOption string
	// 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
	_tripType string
	// 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
	_originLand string
	// 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
	_destination string
	// 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
	_transitLand string
	// 舱位， 用","表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
	_cabin string
	// 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
	_restrictFlightNo string
	// 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
	_excludeFlightNo string
	// 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
	_validDate4Dep string
	// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
	_excludeDateRange4Dep string
	// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripDatePoint4Dep string
	// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_tripExcludeDatePoint4Dep string
	// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
	_flightDateRestrict4Dep string
	// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
	_flightDatePoint4Dep string
	// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
	_saleDate string
	// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
	_adultPassengerIdentity string
	// 最小出行人数,数字1-9
	_minTravelPerson int64
	// 最大出行人数,数字1-9
	_maxTravelPerson int64
	// 小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
	_gv2ChildRule string
	// 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_nationality string
	// 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
	_excludeNationality string
	// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁
	_passengerAge string
	// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
	_ticketPrice int64
	// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
	_adultTax string
	// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
	_childPrice string
	// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
	_childTax string
	// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
	_returnPoint float64
	// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
	_adjustMoney int64
	// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
	_earlyTicketingTimeLimit int64
	// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
	_lateTicketingTimeLimit int64
	// 大客户编码，文本框
	_vipCode string
	// （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
	_fareSource string
	// （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
	_isCreatePnr string
	// 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
	_bookingOffice string
	// 必填项  赋值范围:电子行程单,旅行发票,差额行程单发票,等额行程单
	_receipts string
	// 是否校验票面价,1、可填写 是或者否；默认为否
	_isValidatPrice string
	// （已废除字段）,去程全部未使用可否退票，录入是或否
	_isCanRefund4Dep string
	// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
	_refundPrice4Dep string
	// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
	_refundPartPrice4Dep string
	// （已废除字段）,去程全部未使用可否改期，录入是或否
	_isCanReissue4Dep string
	// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
	_reissuePrice4Dep string
	// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
	_reissuePartPrice4Dep string
	// （已废除字段）,去程NOSHOW规定时限，输入正整数
	_noShowTimeLimit4Dep int64
	// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
	_isNoShowCanRefund4Dep string
	// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
	_isNoShowCanReissue4Dep string
	// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
	_noShowPenalty4Dep int64
	// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
	_luggageRule4Dep string
	// 备注,出票备注文本
	_remark string
	// 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
	_workingHours string
	// （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
	_refundRule string
	// （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
	_reissueRule string
	// （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
	_noshowRule string
	// 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
	_luggageRule string
	// 运价渠道 可选listing,宝贝 默认listing
	_applyChannel string
	// 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
	_commodityType string
	// 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
	_codeSharingType string
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
	_buyTicketNotice string
	// 必填项，全部未使用可否退票，可输入:是，否
	_isCanAllRefund string
	// 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
	_refundFeeAllUnused string
	// 全部未使用退票币种，只能录入币种三字码，默认值CNY
	_refundCurrencyAllUnused string
	// 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_refundFeeTypeAllUnused string
	// 必填项，部分未使用可否退票，可输入：是，否
	_isCanPartRefund string
	// 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
	_refundFeePartUnused string
	// 部分未使用退票币种，可录入币种三字码，默认值CNY
	_refundCurrencyPartUnused string
	// 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_refundFeeTypePartUnused string
	// 必填项，去程可否改期，可输入是或否
	_canDepChange string
	// 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
	_depChangeFee string
	// 去程改期币种，可录入币种三字码，默认值CNY
	_depChangeCurrency string
	// 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
	_depChangeFeeType string
	// 必填项，NOSHOW是否有限制，可输入是或否
	_noshowRestrict string
	// NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
	_noshowTimeRestrict string
	// NOSHOW时限单位(小时/天, 默认为小时)
	_noshowTimeRestrictUnit string
	// NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可退票,不可改期
	_noshowRuleType string
	// NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
	_noshowFee string
	// NOSHOW币种,可录入币种三字码，默认值CNY
	_noshowCurrency string
	// 运价基础，最大长度8
	_farebasis string
	// 运价类型，最大长度3
	_fareTypeCode string
	// 运价tariff，最大长度3
	_tariff string
	// 运价规则id，最大长度4
	_ruleId string
}

// New
