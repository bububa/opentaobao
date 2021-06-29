package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条单程添加 APIRequest
taobao.alitrip.it.fare.addow

自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
*/
type TaobaoAlitripItFareAddowRequest struct {
    model.Params

    // 外部政策ID,1、自行输入的ID，建议为唯一id，有些操作可以使用此id 最多50个字符
    outFileCode   string 

    // 文件编号
    fileCode   string 

    // （后期字段，预留）,产品类型,1.不可为空 2.填写为:包机切位、申请、见舱预订；
    productType   string 

    // （后期字段，预留）,库存模式,1.不可为空 2.填写为见舱或定额；默认为见舱
    stockMode   string 

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

    // 销售日期，1、不得为空 2.输入格式为：2014-04-01~2014-06-30 3.不支持多段组合， 4.3日期格式为 YYYY-MM-DD或YYYY/MM/DD，例：2014-04-01或20104/04/01
    saleDate   string 

    // 成人旅客身份，1.不得为空 2.普通/学生  3.当输入学生时，儿童价格项输入无效 4.当为小团产品时，此适用身份类别必须为 普通。5、后期支持劳工、移民、海员、老人、青年
    adultPassengerIdentity   string 

    // 最小出行人数,数字1-9
    minTravelPerson   int64 

    // 最大出行人数,数字1-9
    maxTravelPerson   int64 

    // 小团儿童计数规则，可选值：1个儿童计1个成人、2个儿童计1个成人、儿童不计
    gv2ChildRule   string 

    // 国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
    nationality   string 

    // 除外国籍，可录入多个用","隔开表示或的关系 1、可录入国家二字代码，为空表示不限制，最多录20个 *默认为空，不输入为不限制
    excludeNationality   string 

    // 乘客年龄,1、可录入范围如21-25表示21周岁至25周岁
    passengerAge   string 

    // 销售票面价,1.不得为空 2.价格区间为【0-999999】 3、销售票面价为10的整数倍(向下取整，如录入3002，则实际录入数值为3000)
    ticketPrice   int64 

    // （后期字段，预留）,成人税费，1、整数金额（包机切位产品适用）
    adultTax   string 

    // 儿童价，1、可不输入，空表示不适用儿童价 2、可输入大于0的正整数及百分比，输入百分比时，成人价格必须录入 例如：2000或70%。 3. 百分比计算的数值，个位向上取整 当"乘客类型"输入非“普通”（成人）时，此项输入无效。
    childPrice   string 

    // （后期字段，预留）,儿童税费，1、整数金额（包机切位产品适用）
    childTax   string 

    // 返点,1.不得为空 2.只允许填写数字，支持到小数点后两位；不用填写% 3.返点需小于100 成人价=销售票面价*（1-返点）+留钱
    returnPoint   float64 

    // 留钱，1.0或正负数字2.-20表示返20元；20代表留20元
    adjustMoney   int64 

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

    // 必填项  赋值范围:电子行程单,旅行发票,差额行程单发票,等额行程单
    receipts   string 

    // 是否校验票面价,1、可填写 是或者否；默认为否
    isValidatPrice   string 

    // （已废除字段）,去程全部未使用可否退票，录入是或否
    isCanRefund4Dep   string 

    // （已废除字段）,去程全部未使用退票费用,可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时退票手续费1000；飞机起飞后不予退票（输入*）；
    refundPrice4Dep   string 

    // （已废除字段）,去程部分未使用退票费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分退票
    refundPartPrice4Dep   string 

    // （已废除字段）,去程全部未使用可否改期，录入是或否
    isCanReissue4Dep   string 

    // （已废除字段）,去程全部未使用改期费用，可输入格式如：200-72-300-48-1000-0-*，表示72小时前改期手续费200；48小时到72小时，改期手续费300；飞机起飞不足48小时改期手续费1000；飞机起飞后不予改期（输入*）；
    reissuePrice4Dep   string 

    // （已废除字段）,去程部分未使用改期费用,可输入空，*或正整数，其中空表示按照航空公司规定执行，*表示不支持部分改期
    reissuePartPrice4Dep   string 

    // （已废除字段）,去程NOSHOW规定时限，输入正整数
    noShowTimeLimit4Dep   int64 

    // （已废除字段）,去程NOSHOW能否退票，输入是或否；默认为否
    isNoShowCanRefund4Dep   string 

    // （已废除字段）,去程NOSHOW能否改期，输入是或否；默认为否
    isNoShowCanReissue4Dep   string 

    // （已废除字段）,去程NOSHOW罚金，可为空，若输入则为正整数；其中空表示按航空公司规定执行
    noShowPenalty4Dep   int64 

    // （后期字段，预留）,去程行李额规定,可输入1-23,1-23 中间用","隔开，表示第一程和第二程（中转）支持行李额为1PC，23KG。若某段为空表示该段按照航空公司规定执行，逗号不可缺少；若不提供免费行李额直接输入空
    luggageRule4Dep   string 

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

    // 运价渠道 可选listing,宝贝 默认listing
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

    // 必填项，NOSHOW是否有限制，可输入是或否
    noshowRestrict   string 

    // NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
    noshowTimeRestrict   string 

    // NOSHOW时限单位(小时/天, 默认为小时)
    noshowTimeRestrictUnit   string 

    // NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期、不可退票,不可改期
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

}

func NewTaobaoAlitripItFareAddowRequest() *TaobaoAlitripItFareAddowRequest{
    return &TaobaoAlitripItFareAddowRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItFareAddowRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.addow"
}

func (r TaobaoAlitripItFareAddowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItFareAddowRequest) SetOutFileCode(outFileCode string) error {
    r.outFileCode = outFileCode
    r.Set("outFileCode", outFileCode)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetOutFileCode() string {
    return r.outFileCode
}

func (r *TaobaoAlitripItFareAddowRequest) SetFileCode(fileCode string) error {
    r.fileCode = fileCode
    r.Set("fileCode", fileCode)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFileCode() string {
    return r.fileCode
}

func (r *TaobaoAlitripItFareAddowRequest) SetProductType(productType string) error {
    r.productType = productType
    r.Set("productType", productType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetProductType() string {
    return r.productType
}

func (r *TaobaoAlitripItFareAddowRequest) SetStockMode(stockMode string) error {
    r.stockMode = stockMode
    r.Set("stockMode", stockMode)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetStockMode() string {
    return r.stockMode
}

func (r *TaobaoAlitripItFareAddowRequest) SetTicketingAirline(ticketingAirline string) error {
    r.ticketingAirline = ticketingAirline
    r.Set("ticketingAirline", ticketingAirline)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTicketingAirline() string {
    return r.ticketingAirline
}

func (r *TaobaoAlitripItFareAddowRequest) SetSaleAirline(saleAirline string) error {
    r.saleAirline = saleAirline
    r.Set("saleAirline", saleAirline)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetSaleAirline() string {
    return r.saleAirline
}

func (r *TaobaoAlitripItFareAddowRequest) SetAddressOption(addressOption string) error {
    r.addressOption = addressOption
    r.Set("addressOption", addressOption)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetAddressOption() string {
    return r.addressOption
}

func (r *TaobaoAlitripItFareAddowRequest) SetTripType(tripType string) error {
    r.tripType = tripType
    r.Set("tripType", tripType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTripType() string {
    return r.tripType
}

func (r *TaobaoAlitripItFareAddowRequest) SetOriginLand(originLand string) error {
    r.originLand = originLand
    r.Set("originLand", originLand)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetOriginLand() string {
    return r.originLand
}

func (r *TaobaoAlitripItFareAddowRequest) SetDestination(destination string) error {
    r.destination = destination
    r.Set("destination", destination)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetDestination() string {
    return r.destination
}

func (r *TaobaoAlitripItFareAddowRequest) SetTransitLand(transitLand string) error {
    r.transitLand = transitLand
    r.Set("transitLand", transitLand)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTransitLand() string {
    return r.transitLand
}

func (r *TaobaoAlitripItFareAddowRequest) SetCabin(cabin string) error {
    r.cabin = cabin
    r.Set("cabin", cabin)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetCabin() string {
    return r.cabin
}

func (r *TaobaoAlitripItFareAddowRequest) SetRestrictFlightNo(restrictFlightNo string) error {
    r.restrictFlightNo = restrictFlightNo
    r.Set("restrictFlightNo", restrictFlightNo)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRestrictFlightNo() string {
    return r.restrictFlightNo
}

func (r *TaobaoAlitripItFareAddowRequest) SetExcludeFlightNo(excludeFlightNo string) error {
    r.excludeFlightNo = excludeFlightNo
    r.Set("excludeFlightNo", excludeFlightNo)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetExcludeFlightNo() string {
    return r.excludeFlightNo
}

func (r *TaobaoAlitripItFareAddowRequest) SetValidDate4Dep(validDate4Dep string) error {
    r.validDate4Dep = validDate4Dep
    r.Set("validDate4Dep", validDate4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetValidDate4Dep() string {
    return r.validDate4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetExcludeDateRange4Dep(excludeDateRange4Dep string) error {
    r.excludeDateRange4Dep = excludeDateRange4Dep
    r.Set("excludeDateRange4Dep", excludeDateRange4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetExcludeDateRange4Dep() string {
    return r.excludeDateRange4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetTripDatePoint4Dep(tripDatePoint4Dep string) error {
    r.tripDatePoint4Dep = tripDatePoint4Dep
    r.Set("tripDatePoint4Dep", tripDatePoint4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTripDatePoint4Dep() string {
    return r.tripDatePoint4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetTripExcludeDatePoint4Dep(tripExcludeDatePoint4Dep string) error {
    r.tripExcludeDatePoint4Dep = tripExcludeDatePoint4Dep
    r.Set("tripExcludeDatePoint4Dep", tripExcludeDatePoint4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTripExcludeDatePoint4Dep() string {
    return r.tripExcludeDatePoint4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetFlightDateRestrict4Dep(flightDateRestrict4Dep string) error {
    r.flightDateRestrict4Dep = flightDateRestrict4Dep
    r.Set("flightDateRestrict4Dep", flightDateRestrict4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFlightDateRestrict4Dep() string {
    return r.flightDateRestrict4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetFlightDatePoint4Dep(flightDatePoint4Dep string) error {
    r.flightDatePoint4Dep = flightDatePoint4Dep
    r.Set("flightDatePoint4Dep", flightDatePoint4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFlightDatePoint4Dep() string {
    return r.flightDatePoint4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetSaleDate(saleDate string) error {
    r.saleDate = saleDate
    r.Set("saleDate", saleDate)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetSaleDate() string {
    return r.saleDate
}

func (r *TaobaoAlitripItFareAddowRequest) SetAdultPassengerIdentity(adultPassengerIdentity string) error {
    r.adultPassengerIdentity = adultPassengerIdentity
    r.Set("adultPassengerIdentity", adultPassengerIdentity)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetAdultPassengerIdentity() string {
    return r.adultPassengerIdentity
}

func (r *TaobaoAlitripItFareAddowRequest) SetMinTravelPerson(minTravelPerson int64) error {
    r.minTravelPerson = minTravelPerson
    r.Set("minTravelPerson", minTravelPerson)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetMinTravelPerson() int64 {
    return r.minTravelPerson
}

func (r *TaobaoAlitripItFareAddowRequest) SetMaxTravelPerson(maxTravelPerson int64) error {
    r.maxTravelPerson = maxTravelPerson
    r.Set("maxTravelPerson", maxTravelPerson)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetMaxTravelPerson() int64 {
    return r.maxTravelPerson
}

func (r *TaobaoAlitripItFareAddowRequest) SetGv2ChildRule(gv2ChildRule string) error {
    r.gv2ChildRule = gv2ChildRule
    r.Set("gv2ChildRule", gv2ChildRule)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetGv2ChildRule() string {
    return r.gv2ChildRule
}

func (r *TaobaoAlitripItFareAddowRequest) SetNationality(nationality string) error {
    r.nationality = nationality
    r.Set("nationality", nationality)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNationality() string {
    return r.nationality
}

func (r *TaobaoAlitripItFareAddowRequest) SetExcludeNationality(excludeNationality string) error {
    r.excludeNationality = excludeNationality
    r.Set("excludeNationality", excludeNationality)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetExcludeNationality() string {
    return r.excludeNationality
}

func (r *TaobaoAlitripItFareAddowRequest) SetPassengerAge(passengerAge string) error {
    r.passengerAge = passengerAge
    r.Set("passengerAge", passengerAge)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetPassengerAge() string {
    return r.passengerAge
}

func (r *TaobaoAlitripItFareAddowRequest) SetTicketPrice(ticketPrice int64) error {
    r.ticketPrice = ticketPrice
    r.Set("ticketPrice", ticketPrice)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTicketPrice() int64 {
    return r.ticketPrice
}

func (r *TaobaoAlitripItFareAddowRequest) SetAdultTax(adultTax string) error {
    r.adultTax = adultTax
    r.Set("adultTax", adultTax)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetAdultTax() string {
    return r.adultTax
}

func (r *TaobaoAlitripItFareAddowRequest) SetChildPrice(childPrice string) error {
    r.childPrice = childPrice
    r.Set("childPrice", childPrice)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetChildPrice() string {
    return r.childPrice
}

func (r *TaobaoAlitripItFareAddowRequest) SetChildTax(childTax string) error {
    r.childTax = childTax
    r.Set("childTax", childTax)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetChildTax() string {
    return r.childTax
}

func (r *TaobaoAlitripItFareAddowRequest) SetReturnPoint(returnPoint float64) error {
    r.returnPoint = returnPoint
    r.Set("returnPoint", returnPoint)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetReturnPoint() float64 {
    return r.returnPoint
}

func (r *TaobaoAlitripItFareAddowRequest) SetAdjustMoney(adjustMoney int64) error {
    r.adjustMoney = adjustMoney
    r.Set("adjustMoney", adjustMoney)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetAdjustMoney() int64 {
    return r.adjustMoney
}

func (r *TaobaoAlitripItFareAddowRequest) SetEarlyTicketingTimeLimit(earlyTicketingTimeLimit int64) error {
    r.earlyTicketingTimeLimit = earlyTicketingTimeLimit
    r.Set("earlyTicketingTimeLimit", earlyTicketingTimeLimit)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetEarlyTicketingTimeLimit() int64 {
    return r.earlyTicketingTimeLimit
}

func (r *TaobaoAlitripItFareAddowRequest) SetLateTicketingTimeLimit(lateTicketingTimeLimit int64) error {
    r.lateTicketingTimeLimit = lateTicketingTimeLimit
    r.Set("lateTicketingTimeLimit", lateTicketingTimeLimit)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetLateTicketingTimeLimit() int64 {
    return r.lateTicketingTimeLimit
}

func (r *TaobaoAlitripItFareAddowRequest) SetVipCode(vipCode string) error {
    r.vipCode = vipCode
    r.Set("vipCode", vipCode)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetVipCode() string {
    return r.vipCode
}

func (r *TaobaoAlitripItFareAddowRequest) SetFareSource(fareSource string) error {
    r.fareSource = fareSource
    r.Set("fareSource", fareSource)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFareSource() string {
    return r.fareSource
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsCreatePnr(isCreatePnr string) error {
    r.isCreatePnr = isCreatePnr
    r.Set("isCreatePnr", isCreatePnr)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsCreatePnr() string {
    return r.isCreatePnr
}

func (r *TaobaoAlitripItFareAddowRequest) SetBookingOffice(bookingOffice string) error {
    r.bookingOffice = bookingOffice
    r.Set("bookingOffice", bookingOffice)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetBookingOffice() string {
    return r.bookingOffice
}

func (r *TaobaoAlitripItFareAddowRequest) SetReceipts(receipts string) error {
    r.receipts = receipts
    r.Set("receipts", receipts)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetReceipts() string {
    return r.receipts
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsValidatPrice(isValidatPrice string) error {
    r.isValidatPrice = isValidatPrice
    r.Set("isValidatPrice", isValidatPrice)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsValidatPrice() string {
    return r.isValidatPrice
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsCanRefund4Dep(isCanRefund4Dep string) error {
    r.isCanRefund4Dep = isCanRefund4Dep
    r.Set("isCanRefund4Dep", isCanRefund4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsCanRefund4Dep() string {
    return r.isCanRefund4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundPrice4Dep(refundPrice4Dep string) error {
    r.refundPrice4Dep = refundPrice4Dep
    r.Set("refundPrice4Dep", refundPrice4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundPrice4Dep() string {
    return r.refundPrice4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundPartPrice4Dep(refundPartPrice4Dep string) error {
    r.refundPartPrice4Dep = refundPartPrice4Dep
    r.Set("refundPartPrice4Dep", refundPartPrice4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundPartPrice4Dep() string {
    return r.refundPartPrice4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsCanReissue4Dep(isCanReissue4Dep string) error {
    r.isCanReissue4Dep = isCanReissue4Dep
    r.Set("isCanReissue4Dep", isCanReissue4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsCanReissue4Dep() string {
    return r.isCanReissue4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetReissuePrice4Dep(reissuePrice4Dep string) error {
    r.reissuePrice4Dep = reissuePrice4Dep
    r.Set("reissuePrice4Dep", reissuePrice4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetReissuePrice4Dep() string {
    return r.reissuePrice4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetReissuePartPrice4Dep(reissuePartPrice4Dep string) error {
    r.reissuePartPrice4Dep = reissuePartPrice4Dep
    r.Set("reissuePartPrice4Dep", reissuePartPrice4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetReissuePartPrice4Dep() string {
    return r.reissuePartPrice4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoShowTimeLimit4Dep(noShowTimeLimit4Dep int64) error {
    r.noShowTimeLimit4Dep = noShowTimeLimit4Dep
    r.Set("noShowTimeLimit4Dep", noShowTimeLimit4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoShowTimeLimit4Dep() int64 {
    return r.noShowTimeLimit4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsNoShowCanRefund4Dep(isNoShowCanRefund4Dep string) error {
    r.isNoShowCanRefund4Dep = isNoShowCanRefund4Dep
    r.Set("isNoShowCanRefund4Dep", isNoShowCanRefund4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsNoShowCanRefund4Dep() string {
    return r.isNoShowCanRefund4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsNoShowCanReissue4Dep(isNoShowCanReissue4Dep string) error {
    r.isNoShowCanReissue4Dep = isNoShowCanReissue4Dep
    r.Set("isNoShowCanReissue4Dep", isNoShowCanReissue4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsNoShowCanReissue4Dep() string {
    return r.isNoShowCanReissue4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoShowPenalty4Dep(noShowPenalty4Dep int64) error {
    r.noShowPenalty4Dep = noShowPenalty4Dep
    r.Set("noShowPenalty4Dep", noShowPenalty4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoShowPenalty4Dep() int64 {
    return r.noShowPenalty4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetLuggageRule4Dep(luggageRule4Dep string) error {
    r.luggageRule4Dep = luggageRule4Dep
    r.Set("luggageRule4Dep", luggageRule4Dep)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetLuggageRule4Dep() string {
    return r.luggageRule4Dep
}

func (r *TaobaoAlitripItFareAddowRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRemark() string {
    return r.remark
}

func (r *TaobaoAlitripItFareAddowRequest) SetWorkingHours(workingHours string) error {
    r.workingHours = workingHours
    r.Set("workingHours", workingHours)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetWorkingHours() string {
    return r.workingHours
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundRule(refundRule string) error {
    r.refundRule = refundRule
    r.Set("refundRule", refundRule)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundRule() string {
    return r.refundRule
}

func (r *TaobaoAlitripItFareAddowRequest) SetReissueRule(reissueRule string) error {
    r.reissueRule = reissueRule
    r.Set("reissueRule", reissueRule)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetReissueRule() string {
    return r.reissueRule
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowRule(noshowRule string) error {
    r.noshowRule = noshowRule
    r.Set("noshowRule", noshowRule)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowRule() string {
    return r.noshowRule
}

func (r *TaobaoAlitripItFareAddowRequest) SetLuggageRule(luggageRule string) error {
    r.luggageRule = luggageRule
    r.Set("luggageRule", luggageRule)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetLuggageRule() string {
    return r.luggageRule
}

func (r *TaobaoAlitripItFareAddowRequest) SetApplyChannel(applyChannel string) error {
    r.applyChannel = applyChannel
    r.Set("applyChannel", applyChannel)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetApplyChannel() string {
    return r.applyChannel
}

func (r *TaobaoAlitripItFareAddowRequest) SetCommodityType(commodityType string) error {
    r.commodityType = commodityType
    r.Set("commodityType", commodityType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetCommodityType() string {
    return r.commodityType
}

func (r *TaobaoAlitripItFareAddowRequest) SetCodeSharingType(codeSharingType string) error {
    r.codeSharingType = codeSharingType
    r.Set("codeSharingType", codeSharingType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetCodeSharingType() string {
    return r.codeSharingType
}

func (r *TaobaoAlitripItFareAddowRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItFareAddowRequest) SetBuyTicketNotice(buyTicketNotice string) error {
    r.buyTicketNotice = buyTicketNotice
    r.Set("buyTicketNotice", buyTicketNotice)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetBuyTicketNotice() string {
    return r.buyTicketNotice
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsCanAllRefund(isCanAllRefund string) error {
    r.isCanAllRefund = isCanAllRefund
    r.Set("isCanAllRefund", isCanAllRefund)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsCanAllRefund() string {
    return r.isCanAllRefund
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundFeeAllUnused(refundFeeAllUnused string) error {
    r.refundFeeAllUnused = refundFeeAllUnused
    r.Set("refundFeeAllUnused", refundFeeAllUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundFeeAllUnused() string {
    return r.refundFeeAllUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundCurrencyAllUnused(refundCurrencyAllUnused string) error {
    r.refundCurrencyAllUnused = refundCurrencyAllUnused
    r.Set("refundCurrencyAllUnused", refundCurrencyAllUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundCurrencyAllUnused() string {
    return r.refundCurrencyAllUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundFeeTypeAllUnused(refundFeeTypeAllUnused string) error {
    r.refundFeeTypeAllUnused = refundFeeTypeAllUnused
    r.Set("refundFeeTypeAllUnused", refundFeeTypeAllUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundFeeTypeAllUnused() string {
    return r.refundFeeTypeAllUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetIsCanPartRefund(isCanPartRefund string) error {
    r.isCanPartRefund = isCanPartRefund
    r.Set("isCanPartRefund", isCanPartRefund)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetIsCanPartRefund() string {
    return r.isCanPartRefund
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundFeePartUnused(refundFeePartUnused string) error {
    r.refundFeePartUnused = refundFeePartUnused
    r.Set("refundFeePartUnused", refundFeePartUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundFeePartUnused() string {
    return r.refundFeePartUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundCurrencyPartUnused(refundCurrencyPartUnused string) error {
    r.refundCurrencyPartUnused = refundCurrencyPartUnused
    r.Set("refundCurrencyPartUnused", refundCurrencyPartUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundCurrencyPartUnused() string {
    return r.refundCurrencyPartUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetRefundFeeTypePartUnused(refundFeeTypePartUnused string) error {
    r.refundFeeTypePartUnused = refundFeeTypePartUnused
    r.Set("refundFeeTypePartUnused", refundFeeTypePartUnused)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRefundFeeTypePartUnused() string {
    return r.refundFeeTypePartUnused
}

func (r *TaobaoAlitripItFareAddowRequest) SetCanDepChange(canDepChange string) error {
    r.canDepChange = canDepChange
    r.Set("canDepChange", canDepChange)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetCanDepChange() string {
    return r.canDepChange
}

func (r *TaobaoAlitripItFareAddowRequest) SetDepChangeFee(depChangeFee string) error {
    r.depChangeFee = depChangeFee
    r.Set("depChangeFee", depChangeFee)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetDepChangeFee() string {
    return r.depChangeFee
}

func (r *TaobaoAlitripItFareAddowRequest) SetDepChangeCurrency(depChangeCurrency string) error {
    r.depChangeCurrency = depChangeCurrency
    r.Set("depChangeCurrency", depChangeCurrency)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetDepChangeCurrency() string {
    return r.depChangeCurrency
}

func (r *TaobaoAlitripItFareAddowRequest) SetDepChangeFeeType(depChangeFeeType string) error {
    r.depChangeFeeType = depChangeFeeType
    r.Set("depChangeFeeType", depChangeFeeType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetDepChangeFeeType() string {
    return r.depChangeFeeType
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowRestrict(noshowRestrict string) error {
    r.noshowRestrict = noshowRestrict
    r.Set("noshowRestrict", noshowRestrict)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowRestrict() string {
    return r.noshowRestrict
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowTimeRestrict(noshowTimeRestrict string) error {
    r.noshowTimeRestrict = noshowTimeRestrict
    r.Set("noshowTimeRestrict", noshowTimeRestrict)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowTimeRestrict() string {
    return r.noshowTimeRestrict
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowTimeRestrictUnit(noshowTimeRestrictUnit string) error {
    r.noshowTimeRestrictUnit = noshowTimeRestrictUnit
    r.Set("noshowTimeRestrictUnit", noshowTimeRestrictUnit)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowTimeRestrictUnit() string {
    return r.noshowTimeRestrictUnit
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowRuleType(noshowRuleType string) error {
    r.noshowRuleType = noshowRuleType
    r.Set("noshowRuleType", noshowRuleType)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowRuleType() string {
    return r.noshowRuleType
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowFee(noshowFee string) error {
    r.noshowFee = noshowFee
    r.Set("noshowFee", noshowFee)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowFee() string {
    return r.noshowFee
}

func (r *TaobaoAlitripItFareAddowRequest) SetNoshowCurrency(noshowCurrency string) error {
    r.noshowCurrency = noshowCurrency
    r.Set("noshowCurrency", noshowCurrency)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetNoshowCurrency() string {
    return r.noshowCurrency
}

func (r *TaobaoAlitripItFareAddowRequest) SetFarebasis(farebasis string) error {
    r.farebasis = farebasis
    r.Set("farebasis", farebasis)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFarebasis() string {
    return r.farebasis
}

func (r *TaobaoAlitripItFareAddowRequest) SetFareTypeCode(fareTypeCode string) error {
    r.fareTypeCode = fareTypeCode
    r.Set("fareTypeCode", fareTypeCode)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetFareTypeCode() string {
    return r.fareTypeCode
}

func (r *TaobaoAlitripItFareAddowRequest) SetTariff(tariff string) error {
    r.tariff = tariff
    r.Set("tariff", tariff)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetTariff() string {
    return r.tariff
}

func (r *TaobaoAlitripItFareAddowRequest) SetRuleId(ruleId string) error {
    r.ruleId = ruleId
    r.Set("ruleId", ruleId)
    return nil
}

func (r TaobaoAlitripItFareAddowRequest) GetRuleId() string {
    return r.ruleId
}

