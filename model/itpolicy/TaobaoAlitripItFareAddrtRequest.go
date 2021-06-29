package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条往返添加 API请求
taobao.alitrip.it.fare.addrt

自有政策往返添加接口
*/
type TaobaoAlitripItFareAddrtRequest struct {
    model.Params
    // 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
    outFileCode   string
    // 文件编号
    fileCode   string
    // （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
    productType   string
    // （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
    stockMode   string
    // 是否1/2RT，1、请填写 是或者否；默认为否
    isRT   string
    // （后期字段，预留）,1/2RT类型，当需要多填入多个时，请以","分隔 1、可填写 、旅行有效期、排除旅行有效期、班期 ；表明1/2RT 混舱计算时，取严还是各取各 2、默认值是 全部各取各
    rtType   string
    // 可组文件编号， 当需要多填入多个时，请以","分隔 1、标记可组文件的编号政策信息，可填写空白； 2、如果是否1/2RT 字段为是，则此字段为必输项
    combinationFilecode   string
    // （后期字段，预留）,是否允许缺口，1、为是或否；默认为否
    isAllowOj   string
    // （后期字段，预留）,缺口类型，1、可填单缺、双缺、始发地缺、目的地缺、或为空；默认为空（当允许缺口组合时，此项为必输项）
    ojType   string
    // （后期字段，预留）,可组缺口文件编号,当需要多填入多个时，请以","分隔 1、标记政策信息，可填写空白； 2、如果是否缺口 字段为是，则此字段为必输项
    combinationOjFilecode   string
    // 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
    ticketingAirline   string
    // 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
    saleAirline   string
    // 城市/机场选项，默认为城市1、可以填写：“机场",“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
    addressOption   string
    // 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
    tripType   string
    // 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
    originLand   string
    // 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
    destination   string
    // 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
    transitLand   string
    // 舱位， 用","表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
    cabin   string
    // 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
    restrictFlightNo   string
    // 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
    excludeFlightNo   string
    // 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
    validDate4Dep   string
    // 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
    excludeDateRange4Dep   string
    // 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    tripDatePoint4Dep   string
    // 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    tripExcludeDatePoint4Dep   string
    // 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
    flightDateRestrict4Dep   string
    // 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    flightDatePoint4Dep   string
    // 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
    validDate4Ret   string
    // 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
    excludeDateRange4Ret   string
    // 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    tripDatePoint4Ret   string
    // 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    tripExcludeDatePoint4Ret   string
    // 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
    flightDateRestrict4Ret   string
    // 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
    flightDatePoint4Ret   string
    // 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
    saleDate   string
    // 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
    minStay   string
    // 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
    maxStay   string
    // 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
    adultPassengerIdentity   string
    // 最小出行人数,数字1-9
    minTravelPerson   int64
    // 最大出行人数,数字1-9
    maxTravelPerson   int64
    // （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
    gv2ChildRule   string
    // 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
    nationality   string
    // 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
    excludeNationality   string
    // 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
    passengerAge   string
    // 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
    ticketPrice   int64
    // （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
    adultTax   int64
    // 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
    childPrice   string
    // （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
    childTax   int64
    // 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
    returnPoint   float64
    // 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
    adjustMoney   int64
    // 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
    rtCommissionFormula   string
    // 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
    earlyTicketingTimeLimit   int64
    // 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
    lateTicketingTimeLimit   int64
    // 大客户编码，文本框
    vipCode   string
    // （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
    fareSource   string
    // （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
    isCreatePnr   string
    // 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
    bookingOffice   string
    // 必填项 赋值范围 境外电子凭证,旅行发票,差额行程单发票,等额行程单
    receipts   string
    // 是否校验票面价,1、可填写 是或者否；默认为否
    isValidatPrice   string
    // （已废除字段）,去程全部未使用可否退票，录入是或否
    isCanRefund4Dep   string
    // （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
    refundPrice4Dep   string
    // （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
    refundPartPrice4Dep   string
    // （已废除字段）,回程全部未使用可否退票，录入是或否
    isCanRefund4Ret   string
    // （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
    refundPrice4Ret   string
    // （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
    refundPartPrice4Ret   string
    // （已废除字段）,去程全部未使用可否改期，录入是或否
    isCanReissue4Dep   string
    // （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
    reissuePrice4Dep   string
    // （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
    reissuePartPrice4Dep   string
    // （已废除字段）,回程全部未使用可否改期，录入是或否
    isCanReissue4Ret   string
    // （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
    reissuePrice4Ret   string
    // （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
    reissuePartPrice4Ret   string
    // （已废除字段）,去程NOSHOW规定时限，输入正整数
    noShowTimeLimit4Dep   int64
    // （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
    isNoShowCanRefund4Dep   string
    // （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
    isNoShowCanReissue4Dep   string
    // （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
    noShowPenalty4Dep   int64
    // （已废除字段）,回程NOSHOW规定时限，输入正整数
    noShowTimeLimit4Ret   int64
    // （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
    isNoShowCanRefund4Ret   string
    // （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
    isNoShowCanReissue4Ret   string
    // （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
    noShowPenalty4Ret   int64
    // （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
    luggageRule4Dep   string
    // （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
    luggageRule4Ret   string
    // 备注,出票备注文本
    remark   string
    // 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
    workingHours   string
    // （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
    refundRule   string
    // （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
    reissueRule   string
    // （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
    noshowRule   string
    // 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
    luggageRule   string
    // 运价渠道 可选listing宝贝  默认listing
    applyChannel   string
    // 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
    commodityType   string
    // 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
    codeSharingType   string
    // json格式的字符串，扩展属性，预留
    extendAttributes   string
    // 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
    buyTicketNotice   string
    // 必填项，全部未使用可否退票，可输入:是，否
    isCanAllRefund   string
    // 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
    refundFeeAllUnused   string
    // 全部未使用退票币种，只能录入币种三字码，默认值CNY
    refundCurrencyAllUnused   string
    // 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    refundFeeTypeAllUnused   string
    // 必填项，部分未使用可否退票，可输入：是，否
    isCanPartRefund   string
    // 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
    refundFeePartUnused   string
    // 部分未使用退票币种，可录入币种三字码，默认值CNY
    refundCurrencyPartUnused   string
    // 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    refundFeeTypePartUnused   string
    // 必填项，去程可否改期，可输入是或否
    canDepChange   string
    // 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
    depChangeFee   string
    // 去程改期币种，可录入币种三字码，默认值CNY
    depChangeCurrency   string
    // 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    depChangeFeeType   string
    // 必填项，回程可否改期，可输入是或否
    canRetChange   string
    // 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
    retChangeFee   string
    // 回程改期币种，可输入币种三字码，默认值CN
    retChangeCurrency   string
    // 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    retChangeFeeType   string
    // 必填项，NOSHOW是否有限制，可输入是或否
    noshowRestrict   string
    // NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
    noshowTimeRestrict   string
    // NOSHOW时限单位(小时/天, 默认为小时)
    noshowTimeRestrictUnit   string
    // NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可改期,不可改期
    noshowRuleType   string
    // NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
    noshowFee   string
    // NOSHOW币种,可录入币种三字码，默认值CNY
    noshowCurrency   string
    // 运价基础，最大长度8
    farebasis   string
    // 运价类型，最大长度3
    fareTypeCode   string
    // 运价tariff，最大长度3
    tariff   string
    // 运价规则id，最大长度4
    ruleId   string
    // 运价组合适用方向,0(或者字段不存在):不限制/1:仅作用在去程/2:仅作用在回程
    fareDirectDestrict   int64
}

// 初始化TaobaoAlitripItFareAddrtRequest对象
func NewTaobaoAlitripItFareAddrtRequest() *TaobaoAlitripItFareAddrtRequest{
    return &TaobaoAlitripItFareAddrtRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareAddrtRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.addrt"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareAddrtRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutFileCode Setter
// 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetOutFileCode(outFileCode string) error {
    r.outFileCode = outFileCode
    r.Set("outFileCode", outFileCode)
    return nil
}

// OutFileCode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetOutFileCode() string {
    return r.outFileCode
}
// FileCode Setter
// 文件编号
func (r *TaobaoAlitripItFareAddrtRequest) SetFileCode(fileCode string) error {
    r.fileCode = fileCode
    r.Set("fileCode", fileCode)
    return nil
}

// FileCode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFileCode() string {
    return r.fileCode
}
// ProductType Setter
// （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
func (r *TaobaoAlitripItFareAddrtRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("productType", productType)
    return nil
}

// ProductType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetProductType() string {
    return r.productType
}
// StockMode Setter
// （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
func (r *TaobaoAlitripItFareAddrtRequest) SetStockMode(stockMode string) error {
    r.stockMode = stockMode
    r.Set("stockMode", stockMode)
    return nil
}

// StockMode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetStockMode() string {
    return r.stockMode
}
// IsRT Setter
// 是否1/2RT，1、请填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsRT(isRT string) error {
    r.isRT = isRT
    r.Set("isRT", isRT)
    return nil
}

// IsRT Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsRT() string {
    return r.isRT
}
// RtType Setter
// （后期字段，预留）,1/2RT类型，当需要多填入多个时，请以","分隔 1、可填写 、旅行有效期、排除旅行有效期、班期 ；表明1/2RT 混舱计算时，取严还是各取各 2、默认值是 全部各取各
func (r *TaobaoAlitripItFareAddrtRequest) SetRtType(rtType string) error {
    r.rtType = rtType
    r.Set("rtType", rtType)
    return nil
}

// RtType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRtType() string {
    return r.rtType
}
// CombinationFilecode Setter
// 可组文件编号， 当需要多填入多个时，请以","分隔 1、标记可组文件的编号政策信息，可填写空白； 2、如果是否1/2RT 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtRequest) SetCombinationFilecode(combinationFilecode string) error {
    r.combinationFilecode = combinationFilecode
    r.Set("combinationFilecode", combinationFilecode)
    return nil
}

// CombinationFilecode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCombinationFilecode() string {
    return r.combinationFilecode
}
// IsAllowOj Setter
// （后期字段，预留）,是否允许缺口，1、为是或否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsAllowOj(isAllowOj string) error {
    r.isAllowOj = isAllowOj
    r.Set("isAllowOj", isAllowOj)
    return nil
}

// IsAllowOj Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsAllowOj() string {
    return r.isAllowOj
}
// OjType Setter
// （后期字段，预留）,缺口类型，1、可填单缺、双缺、始发地缺、目的地缺、或为空；默认为空（当允许缺口组合时，此项为必输项）
func (r *TaobaoAlitripItFareAddrtRequest) SetOjType(ojType string) error {
    r.ojType = ojType
    r.Set("ojType", ojType)
    return nil
}

// OjType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetOjType() string {
    return r.ojType
}
// CombinationOjFilecode Setter
// （后期字段，预留）,可组缺口文件编号,当需要多填入多个时，请以","分隔 1、标记政策信息，可填写空白； 2、如果是否缺口 字段为是，则此字段为必输项
func (r *TaobaoAlitripItFareAddrtRequest) SetCombinationOjFilecode(combinationOjFilecode string) error {
    r.combinationOjFilecode = combinationOjFilecode
    r.Set("combinationOjFilecode", combinationOjFilecode)
    return nil
}

// CombinationOjFilecode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCombinationOjFilecode() string {
    return r.combinationOjFilecode
}
// TicketingAirline Setter
// 出票航司,1.不可为空 2.航空公司二字码 3.只能输入一个
func (r *TaobaoAlitripItFareAddrtRequest) SetTicketingAirline(ticketingAirline string) error {
    r.ticketingAirline = ticketingAirline
    r.Set("ticketingAirline", ticketingAirline)
    return nil
}

// TicketingAirline Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTicketingAirline() string {
    return r.ticketingAirline
}
// SaleAirline Setter
// 销售航司，不同航段之间用 “,”隔开。  1、销售航司二字码； 2、如为直达；请录入一个航司二字码；如为中转，录入格式为 第一程航司，第二程航司；或者航司；若全程都一样，则录入一个航司二字代码即可 3、如果不录入，则航司默认为出票航司；
func (r *TaobaoAlitripItFareAddrtRequest) SetSaleAirline(saleAirline string) error {
    r.saleAirline = saleAirline
    r.Set("saleAirline", saleAirline)
    return nil
}

// SaleAirline Getter
func (r TaobaoAlitripItFareAddrtRequest) GetSaleAirline() string {
    return r.saleAirline
}
// AddressOption Setter
// 城市/机场选项，默认为城市1、可以填写：“机场",“城市”2、定义始发地/目的地/中转点，输入为机场，还是城市。3、如：此项输入机场，则始发地、目的地必须输入机场三字码
func (r *TaobaoAlitripItFareAddrtRequest) SetAddressOption(addressOption string) error {
    r.addressOption = addressOption
    r.Set("addressOption", addressOption)
    return nil
}

// AddressOption Getter
func (r TaobaoAlitripItFareAddrtRequest) GetAddressOption() string {
    return r.addressOption
}
// TripType Setter
// 航程种类，1、默认为直达；有直达和中转两个选项；2、不填写 默认为 直达
func (r *TaobaoAlitripItFareAddrtRequest) SetTripType(tripType string) error {
    r.tripType = tripType
    r.Set("tripType", tripType)
    return nil
}

// TripType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTripType() string {
    return r.tripType
}
// OriginLand Setter
// 始发地,多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtRequest) SetOriginLand(originLand string) error {
    r.originLand = originLand
    r.Set("originLand", originLand)
    return nil
}

// OriginLand Getter
func (r TaobaoAlitripItFareAddrtRequest) GetOriginLand() string {
    return r.originLand
}
// Destination Setter
// 目的地，多个用“,”隔开 1.不得为空 2.可以填写：机场三字码”或“城市码” 3.最多允许100个机场三字码/城市码
func (r *TaobaoAlitripItFareAddrtRequest) SetDestination(destination string) error {
    r.destination = destination
    r.Set("destination", destination)
    return nil
}

// Destination Getter
func (r TaobaoAlitripItFareAddrtRequest) GetDestination() string {
    return r.destination
}
// TransitLand Setter
// 中转地，多个用“，”隔开 1.不得为空 2.可以填写：机场三字码,城市码 3.最多允许100个机场三字码/城市码  4、当航程类型书写为 中转时，此处为必填
func (r *TaobaoAlitripItFareAddrtRequest) SetTransitLand(transitLand string) error {
    r.transitLand = transitLand
    r.Set("transitLand", transitLand)
    return nil
}

// TransitLand Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTransitLand() string {
    return r.transitLand
}
// Cabin Setter
// 舱位， 用","表示航段的分割。 1、舱位代码。每段只允许录入一个舱位代码，若全程舱位一致则可以只录入一个
func (r *TaobaoAlitripItFareAddrtRequest) SetCabin(cabin string) error {
    r.cabin = cabin
    r.Set("cabin", cabin)
    return nil
}

// Cabin Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCabin() string {
    return r.cabin
}
// RestrictFlightNo Setter
// 航班号限制,同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制
func (r *TaobaoAlitripItFareAddrtRequest) SetRestrictFlightNo(restrictFlightNo string) error {
    r.restrictFlightNo = restrictFlightNo
    r.Set("restrictFlightNo", restrictFlightNo)
    return nil
}

// RestrictFlightNo Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRestrictFlightNo() string {
    return r.restrictFlightNo
}
// ExcludeFlightNo Setter
// 排除航班号限制，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班 2 MU  表示所有MU开头的自营航班  3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班 4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班 5 CA(CA)   表示CA自营航班/CA实际承运航班； 6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班； 7 为空表示无限制；8比如两段，第一段无限制，第二段有限制 /CA123
func (r *TaobaoAlitripItFareAddrtRequest) SetExcludeFlightNo(excludeFlightNo string) error {
    r.excludeFlightNo = excludeFlightNo
    r.Set("excludeFlightNo", excludeFlightNo)
    return nil
}

// ExcludeFlightNo Getter
func (r TaobaoAlitripItFareAddrtRequest) GetExcludeFlightNo() string {
    return r.excludeFlightNo
}
// ValidDate4Dep Setter
// 去程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-06-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtRequest) SetValidDate4Dep(validDate4Dep string) error {
    r.validDate4Dep = validDate4Dep
    r.Set("validDate4Dep", validDate4Dep)
    return nil
}

// ValidDate4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetValidDate4Dep() string {
    return r.validDate4Dep
}
// ExcludeDateRange4Dep Setter
// 去程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetExcludeDateRange4Dep(excludeDateRange4Dep string) error {
    r.excludeDateRange4Dep = excludeDateRange4Dep
    r.Set("excludeDateRange4Dep", excludeDateRange4Dep)
    return nil
}

// ExcludeDateRange4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetExcludeDateRange4Dep() string {
    return r.excludeDateRange4Dep
}
// TripDatePoint4Dep Setter
// 去程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetTripDatePoint4Dep(tripDatePoint4Dep string) error {
    r.tripDatePoint4Dep = tripDatePoint4Dep
    r.Set("tripDatePoint4Dep", tripDatePoint4Dep)
    return nil
}

// TripDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTripDatePoint4Dep() string {
    return r.tripDatePoint4Dep
}
// TripExcludeDatePoint4Dep Setter
// 去程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetTripExcludeDatePoint4Dep(tripExcludeDatePoint4Dep string) error {
    r.tripExcludeDatePoint4Dep = tripExcludeDatePoint4Dep
    r.Set("tripExcludeDatePoint4Dep", tripExcludeDatePoint4Dep)
    return nil
}

// TripExcludeDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTripExcludeDatePoint4Dep() string {
    return r.tripExcludeDatePoint4Dep
}
// FlightDateRestrict4Dep Setter
// 去程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtRequest) SetFlightDateRestrict4Dep(flightDateRestrict4Dep string) error {
    r.flightDateRestrict4Dep = flightDateRestrict4Dep
    r.Set("flightDateRestrict4Dep", flightDateRestrict4Dep)
    return nil
}

// FlightDateRestrict4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFlightDateRestrict4Dep() string {
    return r.flightDateRestrict4Dep
}
// FlightDatePoint4Dep Setter
// 去程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetFlightDatePoint4Dep(flightDatePoint4Dep string) error {
    r.flightDatePoint4Dep = flightDatePoint4Dep
    r.Set("flightDatePoint4Dep", flightDatePoint4Dep)
    return nil
}

// FlightDatePoint4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFlightDatePoint4Dep() string {
    return r.flightDatePoint4Dep
}
// ValidDate4Ret Setter
// 回程旅行有效期，支持多段组合，用“,”隔开， 1.不得为空 2例：2014-04-01~2014-6-30，2014-09-01 ~2014-09-30， 3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或2014/04/01
func (r *TaobaoAlitripItFareAddrtRequest) SetValidDate4Ret(validDate4Ret string) error {
    r.validDate4Ret = validDate4Ret
    r.Set("validDate4Ret", validDate4Ret)
    return nil
}

// ValidDate4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetValidDate4Ret() string {
    return r.validDate4Ret
}
// ExcludeDateRange4Ret Setter
// 回程旅行排除时间段，支持多段组合，用“,”隔开隔开， 1.格式，例：2014-04-01~2014-12-31；或例：2014-04-01~2014-06-30,2014-09-01~2014-09-30， 3日期格式为 YYYY-MM-DD,YYYY/MM/DD 4、旅行排除日期最多只能输入200个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetExcludeDateRange4Ret(excludeDateRange4Ret string) error {
    r.excludeDateRange4Ret = excludeDateRange4Ret
    r.Set("excludeDateRange4Ret", excludeDateRange4Ret)
    return nil
}

// ExcludeDateRange4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetExcludeDateRange4Ret() string {
    return r.excludeDateRange4Ret
}
// TripDatePoint4Ret Setter
// 回程旅行日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetTripDatePoint4Ret(tripDatePoint4Ret string) error {
    r.tripDatePoint4Ret = tripDatePoint4Ret
    r.Set("tripDatePoint4Ret", tripDatePoint4Ret)
    return nil
}

// TripDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTripDatePoint4Ret() string {
    return r.tripDatePoint4Ret
}
// TripExcludeDatePoint4Ret Setter
// 回程旅行排除日期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetTripExcludeDatePoint4Ret(tripExcludeDatePoint4Ret string) error {
    r.tripExcludeDatePoint4Ret = tripExcludeDatePoint4Ret
    r.Set("tripExcludeDatePoint4Ret", tripExcludeDatePoint4Ret)
    return nil
}

// TripExcludeDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTripExcludeDatePoint4Ret() string {
    return r.tripExcludeDatePoint4Ret
}
// FlightDateRestrict4Ret Setter
// 回程班期限制，1.12表示周一周二                                                                                              2.12:00-14:00表示每天的12点到14点                                                                                  3. 12:00FRI-12:00SAT 表示周五的中午12点至周六的中午12点
func (r *TaobaoAlitripItFareAddrtRequest) SetFlightDateRestrict4Ret(flightDateRestrict4Ret string) error {
    r.flightDateRestrict4Ret = flightDateRestrict4Ret
    r.Set("flightDateRestrict4Ret", flightDateRestrict4Ret)
    return nil
}

// FlightDateRestrict4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFlightDateRestrict4Ret() string {
    return r.flightDateRestrict4Ret
}
// FlightDatePoint4Ret Setter
// 回程班期作用点，始发航段/第一国际段/主航段/全部；默认空为 第一国际段
func (r *TaobaoAlitripItFareAddrtRequest) SetFlightDatePoint4Ret(flightDatePoint4Ret string) error {
    r.flightDatePoint4Ret = flightDatePoint4Ret
    r.Set("flightDatePoint4Ret", flightDatePoint4Ret)
    return nil
}

// FlightDatePoint4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFlightDatePoint4Ret() string {
    return r.flightDatePoint4Ret
}
// SaleDate Setter
// 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
func (r *TaobaoAlitripItFareAddrtRequest) SetSaleDate(saleDate string) error {
    r.saleDate = saleDate
    r.Set("saleDate", saleDate)
    return nil
}

// SaleDate Getter
func (r TaobaoAlitripItFareAddrtRequest) GetSaleDate() string {
    return r.saleDate
}
// MinStay Setter
// 最短停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtRequest) SetMinStay(minStay string) error {
    r.minStay = minStay
    r.Set("minStay", minStay)
    return nil
}

// MinStay Getter
func (r TaobaoAlitripItFareAddrtRequest) GetMinStay() string {
    return r.minStay
}
// MaxStay Setter
// 最长停留期,1、 默认为空，代表无限制； 2、 格式为：数字+字符/字符 3D表示3天  ; 4M表示4个月 ; SAT表示周六; 3D/SAT表示3天或者周六  3、 12M 表示一年
func (r *TaobaoAlitripItFareAddrtRequest) SetMaxStay(maxStay string) error {
    r.maxStay = maxStay
    r.Set("maxStay", maxStay)
    return nil
}

// MaxStay Getter
func (r TaobaoAlitripItFareAddrtRequest) GetMaxStay() string {
    return r.maxStay
}
// AdultPassengerIdentity Setter
// 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
func (r *TaobaoAlitripItFareAddrtRequest) SetAdultPassengerIdentity(adultPassengerIdentity string) error {
    r.adultPassengerIdentity = adultPassengerIdentity
    r.Set("adultPassengerIdentity", adultPassengerIdentity)
    return nil
}

// AdultPassengerIdentity Getter
func (r TaobaoAlitripItFareAddrtRequest) GetAdultPassengerIdentity() string {
    return r.adultPassengerIdentity
}
// MinTravelPerson Setter
// 最小出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtRequest) SetMinTravelPerson(minTravelPerson int64) error {
    r.minTravelPerson = minTravelPerson
    r.Set("minTravelPerson", minTravelPerson)
    return nil
}

// MinTravelPerson Getter
func (r TaobaoAlitripItFareAddrtRequest) GetMinTravelPerson() int64 {
    return r.minTravelPerson
}
// MaxTravelPerson Setter
// 最大出行人数,数字1-9
func (r *TaobaoAlitripItFareAddrtRequest) SetMaxTravelPerson(maxTravelPerson int64) error {
    r.maxTravelPerson = maxTravelPerson
    r.Set("maxTravelPerson", maxTravelPerson)
    return nil
}

// MaxTravelPerson Getter
func (r TaobaoAlitripItFareAddrtRequest) GetMaxTravelPerson() int64 {
    return r.maxTravelPerson
}
// Gv2ChildRule Setter
// （后期字段，预留）,小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
func (r *TaobaoAlitripItFareAddrtRequest) SetGv2ChildRule(gv2ChildRule string) error {
    r.gv2ChildRule = gv2ChildRule
    r.Set("gv2ChildRule", gv2ChildRule)
    return nil
}

// Gv2ChildRule Getter
func (r TaobaoAlitripItFareAddrtRequest) GetGv2ChildRule() string {
    return r.gv2ChildRule
}
// Nationality Setter
// 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtRequest) SetNationality(nationality string) error {
    r.nationality = nationality
    r.Set("nationality", nationality)
    return nil
}

// Nationality Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNationality() string {
    return r.nationality
}
// ExcludeNationality Setter
// 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
func (r *TaobaoAlitripItFareAddrtRequest) SetExcludeNationality(excludeNationality string) error {
    r.excludeNationality = excludeNationality
    r.Set("excludeNationality", excludeNationality)
    return nil
}

// ExcludeNationality Getter
func (r TaobaoAlitripItFareAddrtRequest) GetExcludeNationality() string {
    return r.excludeNationality
}
// PassengerAge Setter
// 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁,1-表示1岁以上，-99表示99岁以下
func (r *TaobaoAlitripItFareAddrtRequest) SetPassengerAge(passengerAge string) error {
    r.passengerAge = passengerAge
    r.Set("passengerAge", passengerAge)
    return nil
}

// PassengerAge Getter
func (r TaobaoAlitripItFareAddrtRequest) GetPassengerAge() string {
    return r.passengerAge
}
// TicketPrice Setter
// 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
func (r *TaobaoAlitripItFareAddrtRequest) SetTicketPrice(ticketPrice int64) error {
    r.ticketPrice = ticketPrice
    r.Set("ticketPrice", ticketPrice)
    return nil
}

// TicketPrice Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTicketPrice() int64 {
    return r.ticketPrice
}
// AdultTax Setter
// （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtRequest) SetAdultTax(adultTax int64) error {
    r.adultTax = adultTax
    r.Set("adultTax", adultTax)
    return nil
}

// AdultTax Getter
func (r TaobaoAlitripItFareAddrtRequest) GetAdultTax() int64 {
    return r.adultTax
}
// ChildPrice Setter
// 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
func (r *TaobaoAlitripItFareAddrtRequest) SetChildPrice(childPrice string) error {
    r.childPrice = childPrice
    r.Set("childPrice", childPrice)
    return nil
}

// ChildPrice Getter
func (r TaobaoAlitripItFareAddrtRequest) GetChildPrice() string {
    return r.childPrice
}
// ChildTax Setter
// （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
func (r *TaobaoAlitripItFareAddrtRequest) SetChildTax(childTax int64) error {
    r.childTax = childTax
    r.Set("childTax", childTax)
    return nil
}

// ChildTax Getter
func (r TaobaoAlitripItFareAddrtRequest) GetChildTax() int64 {
    return r.childTax
}
// ReturnPoint Setter
// 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
func (r *TaobaoAlitripItFareAddrtRequest) SetReturnPoint(returnPoint float64) error {
    r.returnPoint = returnPoint
    r.Set("returnPoint", returnPoint)
    return nil
}

// ReturnPoint Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReturnPoint() float64 {
    return r.returnPoint
}
// AdjustMoney Setter
// 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
func (r *TaobaoAlitripItFareAddrtRequest) SetAdjustMoney(adjustMoney int64) error {
    r.adjustMoney = adjustMoney
    r.Set("adjustMoney", adjustMoney)
    return nil
}

// AdjustMoney Getter
func (r TaobaoAlitripItFareAddrtRequest) GetAdjustMoney() int64 {
    return r.adjustMoney
}
// RtCommissionFormula Setter
// 1/2RT佣金计算方式,1、各取各，取严； 默认为 取严
func (r *TaobaoAlitripItFareAddrtRequest) SetRtCommissionFormula(rtCommissionFormula string) error {
    r.rtCommissionFormula = rtCommissionFormula
    r.Set("rtCommissionFormula", rtCommissionFormula)
    return nil
}

// RtCommissionFormula Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRtCommissionFormula() string {
    return r.rtCommissionFormula
}
// EarlyTicketingTimeLimit Setter
// 提前出票时限，默认为空，代表无限制； 输入为小于等于365的正整数。 小于或等于最晚出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtRequest) SetEarlyTicketingTimeLimit(earlyTicketingTimeLimit int64) error {
    r.earlyTicketingTimeLimit = earlyTicketingTimeLimit
    r.Set("earlyTicketingTimeLimit", earlyTicketingTimeLimit)
    return nil
}

// EarlyTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtRequest) GetEarlyTicketingTimeLimit() int64 {
    return r.earlyTicketingTimeLimit
}
// LateTicketingTimeLimit Setter
// 最晚出票时限,默认为空，代表无限制； 输入为小于等于365的正整数。 大于或等于提前出票时限。 单位为天
func (r *TaobaoAlitripItFareAddrtRequest) SetLateTicketingTimeLimit(lateTicketingTimeLimit int64) error {
    r.lateTicketingTimeLimit = lateTicketingTimeLimit
    r.Set("lateTicketingTimeLimit", lateTicketingTimeLimit)
    return nil
}

// LateTicketingTimeLimit Getter
func (r TaobaoAlitripItFareAddrtRequest) GetLateTicketingTimeLimit() int64 {
    return r.lateTicketingTimeLimit
}
// VipCode Setter
// 大客户编码，文本框
func (r *TaobaoAlitripItFareAddrtRequest) SetVipCode(vipCode string) error {
    r.vipCode = vipCode
    r.Set("vipCode", vipCode)
    return nil
}

// VipCode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetVipCode() string {
    return r.vipCode
}
// FareSource Setter
// （后期字段，预留）,运价发布渠道,1、可填写 PC、无线、都适用 2、默认为都适用
func (r *TaobaoAlitripItFareAddrtRequest) SetFareSource(fareSource string) error {
    r.fareSource = fareSource
    r.Set("fareSource", fareSource)
    return nil
}

// FareSource Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFareSource() string {
    return r.fareSource
}
// IsCreatePnr Setter
// （后期字段，预留）,是否创建PNR，1、选项 可填写是，否.默认为是
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCreatePnr(isCreatePnr string) error {
    r.isCreatePnr = isCreatePnr
    r.Set("isCreatePnr", isCreatePnr)
    return nil
}

// IsCreatePnr Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCreatePnr() string {
    return r.isCreatePnr
}
// BookingOffice Setter
// 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
func (r *TaobaoAlitripItFareAddrtRequest) SetBookingOffice(bookingOffice string) error {
    r.bookingOffice = bookingOffice
    r.Set("bookingOffice", bookingOffice)
    return nil
}

// BookingOffice Getter
func (r TaobaoAlitripItFareAddrtRequest) GetBookingOffice() string {
    return r.bookingOffice
}
// Receipts Setter
// 必填项 赋值范围 境外电子凭证,旅行发票,差额行程单发票,等额行程单
func (r *TaobaoAlitripItFareAddrtRequest) SetReceipts(receipts string) error {
    r.receipts = receipts
    r.Set("receipts", receipts)
    return nil
}

// Receipts Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReceipts() string {
    return r.receipts
}
// IsValidatPrice Setter
// 是否校验票面价,1、可填写 是或者否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsValidatPrice(isValidatPrice string) error {
    r.isValidatPrice = isValidatPrice
    r.Set("isValidatPrice", isValidatPrice)
    return nil
}

// IsValidatPrice Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsValidatPrice() string {
    return r.isValidatPrice
}
// IsCanRefund4Dep Setter
// （已废除字段）,去程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanRefund4Dep(isCanRefund4Dep string) error {
    r.isCanRefund4Dep = isCanRefund4Dep
    r.Set("isCanRefund4Dep", isCanRefund4Dep)
    return nil
}

// IsCanRefund4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanRefund4Dep() string {
    return r.isCanRefund4Dep
}
// RefundPrice4Dep Setter
// （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundPrice4Dep(refundPrice4Dep string) error {
    r.refundPrice4Dep = refundPrice4Dep
    r.Set("refundPrice4Dep", refundPrice4Dep)
    return nil
}

// RefundPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundPrice4Dep() string {
    return r.refundPrice4Dep
}
// RefundPartPrice4Dep Setter
// （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundPartPrice4Dep(refundPartPrice4Dep string) error {
    r.refundPartPrice4Dep = refundPartPrice4Dep
    r.Set("refundPartPrice4Dep", refundPartPrice4Dep)
    return nil
}

// RefundPartPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundPartPrice4Dep() string {
    return r.refundPartPrice4Dep
}
// IsCanRefund4Ret Setter
// （已废除字段）,回程全部未使用可否退票，录入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanRefund4Ret(isCanRefund4Ret string) error {
    r.isCanRefund4Ret = isCanRefund4Ret
    r.Set("isCanRefund4Ret", isCanRefund4Ret)
    return nil
}

// IsCanRefund4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanRefund4Ret() string {
    return r.isCanRefund4Ret
}
// RefundPrice4Ret Setter
// （已废除字段）,回程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundPrice4Ret(refundPrice4Ret string) error {
    r.refundPrice4Ret = refundPrice4Ret
    r.Set("refundPrice4Ret", refundPrice4Ret)
    return nil
}

// RefundPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundPrice4Ret() string {
    return r.refundPrice4Ret
}
// RefundPartPrice4Ret Setter
// （已废除字段）,回程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundPartPrice4Ret(refundPartPrice4Ret string) error {
    r.refundPartPrice4Ret = refundPartPrice4Ret
    r.Set("refundPartPrice4Ret", refundPartPrice4Ret)
    return nil
}

// RefundPartPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundPartPrice4Ret() string {
    return r.refundPartPrice4Ret
}
// IsCanReissue4Dep Setter
// （已废除字段）,去程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanReissue4Dep(isCanReissue4Dep string) error {
    r.isCanReissue4Dep = isCanReissue4Dep
    r.Set("isCanReissue4Dep", isCanReissue4Dep)
    return nil
}

// IsCanReissue4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanReissue4Dep() string {
    return r.isCanReissue4Dep
}
// ReissuePrice4Dep Setter
// （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtRequest) SetReissuePrice4Dep(reissuePrice4Dep string) error {
    r.reissuePrice4Dep = reissuePrice4Dep
    r.Set("reissuePrice4Dep", reissuePrice4Dep)
    return nil
}

// ReissuePrice4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReissuePrice4Dep() string {
    return r.reissuePrice4Dep
}
// ReissuePartPrice4Dep Setter
// （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtRequest) SetReissuePartPrice4Dep(reissuePartPrice4Dep string) error {
    r.reissuePartPrice4Dep = reissuePartPrice4Dep
    r.Set("reissuePartPrice4Dep", reissuePartPrice4Dep)
    return nil
}

// ReissuePartPrice4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReissuePartPrice4Dep() string {
    return r.reissuePartPrice4Dep
}
// IsCanReissue4Ret Setter
// （已废除字段）,回程全部未使用可否改期，录入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanReissue4Ret(isCanReissue4Ret string) error {
    r.isCanReissue4Ret = isCanReissue4Ret
    r.Set("isCanReissue4Ret", isCanReissue4Ret)
    return nil
}

// IsCanReissue4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanReissue4Ret() string {
    return r.isCanReissue4Ret
}
// ReissuePrice4Ret Setter
// （已废除字段）,回程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
func (r *TaobaoAlitripItFareAddrtRequest) SetReissuePrice4Ret(reissuePrice4Ret string) error {
    r.reissuePrice4Ret = reissuePrice4Ret
    r.Set("reissuePrice4Ret", reissuePrice4Ret)
    return nil
}

// ReissuePrice4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReissuePrice4Ret() string {
    return r.reissuePrice4Ret
}
// ReissuePartPrice4Ret Setter
// （已废除字段）,回程部分未使用改期费用，可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
func (r *TaobaoAlitripItFareAddrtRequest) SetReissuePartPrice4Ret(reissuePartPrice4Ret string) error {
    r.reissuePartPrice4Ret = reissuePartPrice4Ret
    r.Set("reissuePartPrice4Ret", reissuePartPrice4Ret)
    return nil
}

// ReissuePartPrice4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReissuePartPrice4Ret() string {
    return r.reissuePartPrice4Ret
}
// NoShowTimeLimit4Dep Setter
// （已废除字段）,去程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtRequest) SetNoShowTimeLimit4Dep(noShowTimeLimit4Dep int64) error {
    r.noShowTimeLimit4Dep = noShowTimeLimit4Dep
    r.Set("noShowTimeLimit4Dep", noShowTimeLimit4Dep)
    return nil
}

// NoShowTimeLimit4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoShowTimeLimit4Dep() int64 {
    return r.noShowTimeLimit4Dep
}
// IsNoShowCanRefund4Dep Setter
// （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsNoShowCanRefund4Dep(isNoShowCanRefund4Dep string) error {
    r.isNoShowCanRefund4Dep = isNoShowCanRefund4Dep
    r.Set("isNoShowCanRefund4Dep", isNoShowCanRefund4Dep)
    return nil
}

// IsNoShowCanRefund4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsNoShowCanRefund4Dep() string {
    return r.isNoShowCanRefund4Dep
}
// IsNoShowCanReissue4Dep Setter
// （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsNoShowCanReissue4Dep(isNoShowCanReissue4Dep string) error {
    r.isNoShowCanReissue4Dep = isNoShowCanReissue4Dep
    r.Set("isNoShowCanReissue4Dep", isNoShowCanReissue4Dep)
    return nil
}

// IsNoShowCanReissue4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsNoShowCanReissue4Dep() string {
    return r.isNoShowCanReissue4Dep
}
// NoShowPenalty4Dep Setter
// （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtRequest) SetNoShowPenalty4Dep(noShowPenalty4Dep int64) error {
    r.noShowPenalty4Dep = noShowPenalty4Dep
    r.Set("noShowPenalty4Dep", noShowPenalty4Dep)
    return nil
}

// NoShowPenalty4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoShowPenalty4Dep() int64 {
    return r.noShowPenalty4Dep
}
// NoShowTimeLimit4Ret Setter
// （已废除字段）,回程NOSHOW规定时限，输入正整数
func (r *TaobaoAlitripItFareAddrtRequest) SetNoShowTimeLimit4Ret(noShowTimeLimit4Ret int64) error {
    r.noShowTimeLimit4Ret = noShowTimeLimit4Ret
    r.Set("noShowTimeLimit4Ret", noShowTimeLimit4Ret)
    return nil
}

// NoShowTimeLimit4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoShowTimeLimit4Ret() int64 {
    return r.noShowTimeLimit4Ret
}
// IsNoShowCanRefund4Ret Setter
// （已废除字段）,回程NOSHOW能否退票，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsNoShowCanRefund4Ret(isNoShowCanRefund4Ret string) error {
    r.isNoShowCanRefund4Ret = isNoShowCanRefund4Ret
    r.Set("isNoShowCanRefund4Ret", isNoShowCanRefund4Ret)
    return nil
}

// IsNoShowCanRefund4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsNoShowCanRefund4Ret() string {
    return r.isNoShowCanRefund4Ret
}
// IsNoShowCanReissue4Ret Setter
// （已废除字段）,回程NOSHOW能否改期，输入是或否；默认为否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsNoShowCanReissue4Ret(isNoShowCanReissue4Ret string) error {
    r.isNoShowCanReissue4Ret = isNoShowCanReissue4Ret
    r.Set("isNoShowCanReissue4Ret", isNoShowCanReissue4Ret)
    return nil
}

// IsNoShowCanReissue4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsNoShowCanReissue4Ret() string {
    return r.isNoShowCanReissue4Ret
}
// NoShowPenalty4Ret Setter
// （已废除字段）,回程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
func (r *TaobaoAlitripItFareAddrtRequest) SetNoShowPenalty4Ret(noShowPenalty4Ret int64) error {
    r.noShowPenalty4Ret = noShowPenalty4Ret
    r.Set("noShowPenalty4Ret", noShowPenalty4Ret)
    return nil
}

// NoShowPenalty4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoShowPenalty4Ret() int64 {
    return r.noShowPenalty4Ret
}
// LuggageRule4Dep Setter
// （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtRequest) SetLuggageRule4Dep(luggageRule4Dep string) error {
    r.luggageRule4Dep = luggageRule4Dep
    r.Set("luggageRule4Dep", luggageRule4Dep)
    return nil
}

// LuggageRule4Dep Getter
func (r TaobaoAlitripItFareAddrtRequest) GetLuggageRule4Dep() string {
    return r.luggageRule4Dep
}
// LuggageRule4Ret Setter
// （后期字段，预留）,回程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
func (r *TaobaoAlitripItFareAddrtRequest) SetLuggageRule4Ret(luggageRule4Ret string) error {
    r.luggageRule4Ret = luggageRule4Ret
    r.Set("luggageRule4Ret", luggageRule4Ret)
    return nil
}

// LuggageRule4Ret Getter
func (r TaobaoAlitripItFareAddrtRequest) GetLuggageRule4Ret() string {
    return r.luggageRule4Ret
}
// Remark Setter
// 备注,出票备注文本
func (r *TaobaoAlitripItFareAddrtRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRemark() string {
    return r.remark
}
// WorkingHours Setter
// 工作时间,18:00FRI表示周一到周五的每天早上9点到下午6点                                                     最多录入三个时间段用，隔开表示或的关系                                                                               可以为空，表示不限制(运价上的工作时间优先级高于设置时间界面上的时间)
func (r *TaobaoAlitripItFareAddrtRequest) SetWorkingHours(workingHours string) error {
    r.workingHours = workingHours
    r.Set("workingHours", workingHours)
    return nil
}

// WorkingHours Getter
func (r TaobaoAlitripItFareAddrtRequest) GetWorkingHours() string {
    return r.workingHours
}
// RefundRule Setter
// （已废除字段）退票规定,1、不可为空 2、可填写：收取20%退票费用，或者是收取500元退票费等。 3、退票规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundRule(refundRule string) error {
    r.refundRule = refundRule
    r.Set("refundRule", refundRule)
    return nil
}

// RefundRule Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundRule() string {
    return r.refundRule
}
// ReissueRule Setter
// （已废除字段）改期规定,1、不可为空 2、可填写：收取20%改期费用，或者是收取500元改期费等。 3、改期规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetReissueRule(reissueRule string) error {
    r.reissueRule = reissueRule
    r.Set("reissueRule", reissueRule)
    return nil
}

// ReissueRule Getter
func (r TaobaoAlitripItFareAddrtRequest) GetReissueRule() string {
    return r.reissueRule
}
// NoshowRule Setter
// （已废除字段）误机罚金说明，1、不可为空 2、可填写：起飞前不得退票，不得改期 3、误机罚金说明最多为300个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowRule(noshowRule string) error {
    r.noshowRule = noshowRule
    r.Set("noshowRule", noshowRule)
    return nil
}

// NoshowRule Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowRule() string {
    return r.noshowRule
}
// LuggageRule Setter
// 行李额规定,1、不可为空2、可填写：1PC。逾重行李费用为每公斤100元3、行李额规定最多为300个字符
func (r *TaobaoAlitripItFareAddrtRequest) SetLuggageRule(luggageRule string) error {
    r.luggageRule = luggageRule
    r.Set("luggageRule", luggageRule)
    return nil
}

// LuggageRule Getter
func (r TaobaoAlitripItFareAddrtRequest) GetLuggageRule() string {
    return r.luggageRule
}
// ApplyChannel Setter
// 运价渠道 可选listing宝贝  默认listing
func (r *TaobaoAlitripItFareAddrtRequest) SetApplyChannel(applyChannel string) error {
    r.applyChannel = applyChannel
    r.Set("applyChannel", applyChannel)
    return nil
}

// ApplyChannel Getter
func (r TaobaoAlitripItFareAddrtRequest) GetApplyChannel() string {
    return r.applyChannel
}
// CommodityType Setter
// 商品类型，可选值：普通、金牌，默认普通，非金牌卖家不得选择金牌
func (r *TaobaoAlitripItFareAddrtRequest) SetCommodityType(commodityType string) error {
    r.commodityType = commodityType
    r.Set("commodityType", commodityType)
    return nil
}

// CommodityType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCommodityType() string {
    return r.commodityType
}
// CodeSharingType Setter
// 不录入表示不限制；选项为：仅限同集团代码共享适用；代码共享适用；不允许代码共享；不限制 默认不限制
func (r *TaobaoAlitripItFareAddrtRequest) SetCodeSharingType(codeSharingType string) error {
    r.codeSharingType = codeSharingType
    r.Set("codeSharingType", codeSharingType)
    return nil
}

// CodeSharingType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCodeSharingType() string {
    return r.codeSharingType
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareAddrtRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareAddrtRequest) GetExtendAttributes() string {
    return r.extendAttributes
}
// BuyTicketNotice Setter
// 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
func (r *TaobaoAlitripItFareAddrtRequest) SetBuyTicketNotice(buyTicketNotice string) error {
    r.buyTicketNotice = buyTicketNotice
    r.Set("buyTicketNotice", buyTicketNotice)
    return nil
}

// BuyTicketNotice Getter
func (r TaobaoAlitripItFareAddrtRequest) GetBuyTicketNotice() string {
    return r.buyTicketNotice
}
// IsCanAllRefund Setter
// 必填项，全部未使用可否退票，可输入:是，否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanAllRefund(isCanAllRefund string) error {
    r.isCanAllRefund = isCanAllRefund
    r.Set("isCanAllRefund", isCanAllRefund)
    return nil
}

// IsCanAllRefund Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanAllRefund() string {
    return r.isCanAllRefund
}
// RefundFeeAllUnused Setter
// 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如:  1) 200 表示退票手续费为200（货币单位在下一个格子里） 2) 20% 表示退票手续费为票面价的20% 3）* 表示不允许退票 4) 200-0-400 表示起飞前退票手续费200；起飞后退票手续费400 5) 30%-0-* 表示起飞前退票手续费为票面价的30%；起飞后不允许退票 6）200-72-300-48-1000-0-* 表示72小时前退票手续费200; 48小时到72小时，退票手续费300; 飞机起飞不足48小时; 退票手续费1000; 飞机起飞后不予退票（输入*） 7) 10%-72-30%-48-70%-0-* 表示72小时前退票手续费为票面价的10%; 48小时到72小时，退票手续费为票面价的30%; 飞机起飞不足48小时; 退票手续费为票面价的70%; 飞机起飞后不予退票（输入*）
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundFeeAllUnused(refundFeeAllUnused string) error {
    r.refundFeeAllUnused = refundFeeAllUnused
    r.Set("refundFeeAllUnused", refundFeeAllUnused)
    return nil
}

// RefundFeeAllUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundFeeAllUnused() string {
    return r.refundFeeAllUnused
}
// RefundCurrencyAllUnused Setter
// 全部未使用退票币种，只能录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundCurrencyAllUnused(refundCurrencyAllUnused string) error {
    r.refundCurrencyAllUnused = refundCurrencyAllUnused
    r.Set("refundCurrencyAllUnused", refundCurrencyAllUnused)
    return nil
}

// RefundCurrencyAllUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundCurrencyAllUnused() string {
    return r.refundCurrencyAllUnused
}
// RefundFeeTypeAllUnused Setter
// 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundFeeTypeAllUnused(refundFeeTypeAllUnused string) error {
    r.refundFeeTypeAllUnused = refundFeeTypeAllUnused
    r.Set("refundFeeTypeAllUnused", refundFeeTypeAllUnused)
    return nil
}

// RefundFeeTypeAllUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundFeeTypeAllUnused() string {
    return r.refundFeeTypeAllUnused
}
// IsCanPartRefund Setter
// 必填项，部分未使用可否退票，可输入：是，否
func (r *TaobaoAlitripItFareAddrtRequest) SetIsCanPartRefund(isCanPartRefund string) error {
    r.isCanPartRefund = isCanPartRefund
    r.Set("isCanPartRefund", isCanPartRefund)
    return nil
}

// IsCanPartRefund Getter
func (r TaobaoAlitripItFareAddrtRequest) GetIsCanPartRefund() string {
    return r.isCanPartRefund
}
// RefundFeePartUnused Setter
// 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundFeePartUnused(refundFeePartUnused string) error {
    r.refundFeePartUnused = refundFeePartUnused
    r.Set("refundFeePartUnused", refundFeePartUnused)
    return nil
}

// RefundFeePartUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundFeePartUnused() string {
    return r.refundFeePartUnused
}
// RefundCurrencyPartUnused Setter
// 部分未使用退票币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundCurrencyPartUnused(refundCurrencyPartUnused string) error {
    r.refundCurrencyPartUnused = refundCurrencyPartUnused
    r.Set("refundCurrencyPartUnused", refundCurrencyPartUnused)
    return nil
}

// RefundCurrencyPartUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundCurrencyPartUnused() string {
    return r.refundCurrencyPartUnused
}
// RefundFeeTypePartUnused Setter
// 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtRequest) SetRefundFeeTypePartUnused(refundFeeTypePartUnused string) error {
    r.refundFeeTypePartUnused = refundFeeTypePartUnused
    r.Set("refundFeeTypePartUnused", refundFeeTypePartUnused)
    return nil
}

// RefundFeeTypePartUnused Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRefundFeeTypePartUnused() string {
    return r.refundFeeTypePartUnused
}
// CanDepChange Setter
// 必填项，去程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetCanDepChange(canDepChange string) error {
    r.canDepChange = canDepChange
    r.Set("canDepChange", canDepChange)
    return nil
}

// CanDepChange Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCanDepChange() string {
    return r.canDepChange
}
// DepChangeFee Setter
// 【去程可否改期】为是时为必填项， 可输入格式如:  1) 200 表示改期手续费为200（货币单位在下一个格子里） 2）* 表示不允许改期 3) 200-0-400 表示起飞前改期手续费200；起飞后改期手续费400 4) 30-0-* 表示起飞前改期手续费30；起飞后不允许改期 5）200-72-300-48-1000-0-* 表示72小时前改期手续费200; 48小时到72小时，改期手续费300; 飞机起飞不足48小时; 改期手续费1000; 飞机起飞后不予改期（输入*）
func (r *TaobaoAlitripItFareAddrtRequest) SetDepChangeFee(depChangeFee string) error {
    r.depChangeFee = depChangeFee
    r.Set("depChangeFee", depChangeFee)
    return nil
}

// DepChangeFee Getter
func (r TaobaoAlitripItFareAddrtRequest) GetDepChangeFee() string {
    return r.depChangeFee
}
// DepChangeCurrency Setter
// 去程改期币种，可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtRequest) SetDepChangeCurrency(depChangeCurrency string) error {
    r.depChangeCurrency = depChangeCurrency
    r.Set("depChangeCurrency", depChangeCurrency)
    return nil
}

// DepChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtRequest) GetDepChangeCurrency() string {
    return r.depChangeCurrency
}
// DepChangeFeeType Setter
// 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtRequest) SetDepChangeFeeType(depChangeFeeType string) error {
    r.depChangeFeeType = depChangeFeeType
    r.Set("depChangeFeeType", depChangeFeeType)
    return nil
}

// DepChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetDepChangeFeeType() string {
    return r.depChangeFeeType
}
// CanRetChange Setter
// 必填项，回程可否改期，可输入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetCanRetChange(canRetChange string) error {
    r.canRetChange = canRetChange
    r.Set("canRetChange", canRetChange)
    return nil
}

// CanRetChange Getter
func (r TaobaoAlitripItFareAddrtRequest) GetCanRetChange() string {
    return r.canRetChange
}
// RetChangeFee Setter
// 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
func (r *TaobaoAlitripItFareAddrtRequest) SetRetChangeFee(retChangeFee string) error {
    r.retChangeFee = retChangeFee
    r.Set("retChangeFee", retChangeFee)
    return nil
}

// RetChangeFee Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRetChangeFee() string {
    return r.retChangeFee
}
// RetChangeCurrency Setter
// 回程改期币种，可输入币种三字码，默认值CN
func (r *TaobaoAlitripItFareAddrtRequest) SetRetChangeCurrency(retChangeCurrency string) error {
    r.retChangeCurrency = retChangeCurrency
    r.Set("retChangeCurrency", retChangeCurrency)
    return nil
}

// RetChangeCurrency Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRetChangeCurrency() string {
    return r.retChangeCurrency
}
// RetChangeFeeType Setter
// 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
func (r *TaobaoAlitripItFareAddrtRequest) SetRetChangeFeeType(retChangeFeeType string) error {
    r.retChangeFeeType = retChangeFeeType
    r.Set("retChangeFeeType", retChangeFeeType)
    return nil
}

// RetChangeFeeType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRetChangeFeeType() string {
    return r.retChangeFeeType
}
// NoshowRestrict Setter
// 必填项，NOSHOW是否有限制，可输入是或否
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowRestrict(noshowRestrict string) error {
    r.noshowRestrict = noshowRestrict
    r.Set("noshowRestrict", noshowRestrict)
    return nil
}

// NoshowRestrict Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowRestrict() string {
    return r.noshowRestrict
}
// NoshowTimeRestrict Setter
// NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowTimeRestrict(noshowTimeRestrict string) error {
    r.noshowTimeRestrict = noshowTimeRestrict
    r.Set("noshowTimeRestrict", noshowTimeRestrict)
    return nil
}

// NoshowTimeRestrict Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowTimeRestrict() string {
    return r.noshowTimeRestrict
}
// NoshowTimeRestrictUnit Setter
// NOSHOW时限单位(小时/天, 默认为小时)
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowTimeRestrictUnit(noshowTimeRestrictUnit string) error {
    r.noshowTimeRestrictUnit = noshowTimeRestrictUnit
    r.Set("noshowTimeRestrictUnit", noshowTimeRestrictUnit)
    return nil
}

// NoshowTimeRestrictUnit Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowTimeRestrictUnit() string {
    return r.noshowTimeRestrictUnit
}
// NoshowRuleType Setter
// NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可改期,不可改期
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowRuleType(noshowRuleType string) error {
    r.noshowRuleType = noshowRuleType
    r.Set("noshowRuleType", noshowRuleType)
    return nil
}

// NoshowRuleType Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowRuleType() string {
    return r.noshowRuleType
}
// NoshowFee Setter
// NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowFee(noshowFee string) error {
    r.noshowFee = noshowFee
    r.Set("noshowFee", noshowFee)
    return nil
}

// NoshowFee Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowFee() string {
    return r.noshowFee
}
// NoshowCurrency Setter
// NOSHOW币种,可录入币种三字码，默认值CNY
func (r *TaobaoAlitripItFareAddrtRequest) SetNoshowCurrency(noshowCurrency string) error {
    r.noshowCurrency = noshowCurrency
    r.Set("noshowCurrency", noshowCurrency)
    return nil
}

// NoshowCurrency Getter
func (r TaobaoAlitripItFareAddrtRequest) GetNoshowCurrency() string {
    return r.noshowCurrency
}
// Farebasis Setter
// 运价基础，最大长度8
func (r *TaobaoAlitripItFareAddrtRequest) SetFarebasis(farebasis string) error {
    r.farebasis = farebasis
    r.Set("farebasis", farebasis)
    return nil
}

// Farebasis Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFarebasis() string {
    return r.farebasis
}
// FareTypeCode Setter
// 运价类型，最大长度3
func (r *TaobaoAlitripItFareAddrtRequest) SetFareTypeCode(fareTypeCode string) error {
    r.fareTypeCode = fareTypeCode
    r.Set("fareTypeCode", fareTypeCode)
    return nil
}

// FareTypeCode Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFareTypeCode() string {
    return r.fareTypeCode
}
// Tariff Setter
// 运价tariff，最大长度3
func (r *TaobaoAlitripItFareAddrtRequest) SetTariff(tariff string) error {
    r.tariff = tariff
    r.Set("tariff", tariff)
    return nil
}

// Tariff Getter
func (r TaobaoAlitripItFareAddrtRequest) GetTariff() string {
    return r.tariff
}
// RuleId Setter
// 运价规则id，最大长度4
func (r *TaobaoAlitripItFareAddrtRequest) SetRuleId(ruleId string) error {
    r.ruleId = ruleId
    r.Set("ruleId", ruleId)
    return nil
}

// RuleId Getter
func (r TaobaoAlitripItFareAddrtRequest) GetRuleId() string {
    return r.ruleId
}
// FareDirectDestrict Setter
// 运价组合适用方向,0(或者字段不存在):不限制/1:仅作用在去程/2:仅作用在回程
func (r *TaobaoAlitripItFareAddrtRequest) SetFareDirectDestrict(fareDirectDestrict int64) error {
    r.fareDirectDestrict = fareDirectDestrict
    r.Set("fareDirectDestrict", fareDirectDestrict)
    return nil
}

// FareDirectDestrict Getter
func (r TaobaoAlitripItFareAddrtRequest) GetFareDirectDestrict() int64 {
    return r.fareDirectDestrict
}
