package flight

import (
	"sync"
)

// PolicyDto 结构体
type PolicyDto struct {
	// 行程信息
	Flights []FlightDto `json:"flights,omitempty" xml:"flights>flight_dto,omitempty"`
	// 排除航线：格式为：KKK-*，*-KKK，KKK-XXX三种格式；三字码，大写；*表示全部；最多允许4000字符；多个请用英文,隔开
	NotApplOd []string `json:"not_appl_od,omitempty" xml:"not_appl_od>string,omitempty"`
	// 运价类型：0，FD；1，NFD；2，特殊运价；5，IBE；11，B2T；12，旗舰店；13，官网；14，大卖家
	FareSource []string `json:"fare_source,omitempty" xml:"fare_source>string,omitempty"`
	// 航司二字码
	Airline string `json:"airline,omitempty" xml:"airline,omitempty"`
	// 到达机场三字码，不填代表不限，不要填ALL
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 出发机场三字码，不填代表不限，不要填ALL
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 商家配置号
	OfficeNo string `json:"office_no,omitempty" xml:"office_no,omitempty"`
	// 政策编码
	PolicyCode string `json:"policy_code,omitempty" xml:"policy_code,omitempty"`
	// 政策备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 字符；最多200；可输入多个航司二字码,英文逗号分隔；航司二字码
	CodeShareAirline string `json:"code_share_airline,omitempty" xml:"code_share_airline,omitempty"`
	// 大客户编码，请输入包含PAT在内的完整格式。
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
	// 用于匹配平台退改，如果填写，则只能填写精确的farebasis，不得带通配符；匹配不到，走平台默认退改
	FareBasis string `json:"fare_basis,omitempty" xml:"fare_basis,omitempty"`
	// 中转点
	TransitAirport string `json:"transit_airport,omitempty" xml:"transit_airport,omitempty"`
	// 前段可衔接航司
	BeforeCombineAirline string `json:"before_combine_airline,omitempty" xml:"before_combine_airline,omitempty"`
	// 后段可衔接航司
	NextCombineAirline string `json:"next_combine_airline,omitempty" xml:"next_combine_airline,omitempty"`
	// 后段可衔接城市(&#34;,&#34;拼接)
	NextCombineCity string `json:"next_combine_city,omitempty" xml:"next_combine_city,omitempty"`
	// 是否订位：1，平台订位；0，平台不订位；2，紧张订位
	CreatePnr int64 `json:"create_pnr,omitempty" xml:"create_pnr,omitempty"`
	// 政策来源：0，手工政策；1，excel政策；2，api政策
	PolicySource int64 `json:"policy_source,omitempty" xml:"policy_source,omitempty"`
	// 价格信息栏
	Price *PriceDto `json:"price,omitempty" xml:"price,omitempty"`
	// 销售信息
	Sale *SaleDto `json:"sale,omitempty" xml:"sale,omitempty"`
	// 政策状态：1，有效；2，挂起；0，删除 3，停用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 0-白名单，航班舱位命中填写的去程/回程舱位，则可以使用本条政策；1-黑名单，航班舱位命中填写的去程/回程舱位，则不可以使用本条政策；默认为0-白名单。
	IsWhite int64 `json:"is_white,omitempty" xml:"is_white,omitempty"`
	// 是否验价：0，不验价；1，验价
	Pata int64 `json:"pata,omitempty" xml:"pata,omitempty"`
	// 共享航班：0，不支持代码共享；1，支持代码共享；2，仅支持代码共享
	CodeShare int64 `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// 直达中转：0，直达；1，中转；默认为0
	DirectTransferType int64 `json:"direct_transfer_type,omitempty" xml:"direct_transfer_type,omitempty"`
	// 乘客限制
	Passenger *PassengerDto `json:"passenger,omitempty" xml:"passenger,omitempty"`
	// 外放舱位数量小于等于阈值时订位或停售
	CreatePnrLimit int64 `json:"create_pnr_limit,omitempty" xml:"create_pnr_limit,omitempty"`
	// 指定票面价，单位为元
	FarePrice int64 `json:"fare_price,omitempty" xml:"fare_price,omitempty"`
	// 0-往返两程总价， 1-去回程分别定价， 默认为0
	PriceType int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 是否支持共享出主飞;是:true，否:false
	OutMainFlight bool `json:"out_main_flight,omitempty" xml:"out_main_flight,omitempty"`
	// 是否中转运价.true:是，false:否
	IsOwEoe bool `json:"is_ow_eoe,omitempty" xml:"is_ow_eoe,omitempty"`
}

var poolPolicyDto = sync.Pool{
	New: func() any {
		return new(PolicyDto)
	},
}

// GetPolicyDto() 从对象池中获取PolicyDto
func GetPolicyDto() *PolicyDto {
	return poolPolicyDto.Get().(*PolicyDto)
}

// ReleasePolicyDto 释放PolicyDto
func ReleasePolicyDto(v *PolicyDto) {
	v.Flights = v.Flights[:0]
	v.NotApplOd = v.NotApplOd[:0]
	v.FareSource = v.FareSource[:0]
	v.Airline = ""
	v.ArrAirport = ""
	v.DepAirport = ""
	v.OfficeNo = ""
	v.PolicyCode = ""
	v.Memo = ""
	v.CodeShareAirline = ""
	v.AccountCode = ""
	v.FareBasis = ""
	v.TransitAirport = ""
	v.BeforeCombineAirline = ""
	v.NextCombineAirline = ""
	v.NextCombineCity = ""
	v.CreatePnr = 0
	v.PolicySource = 0
	v.Price = nil
	v.Sale = nil
	v.Status = 0
	v.IsWhite = 0
	v.Pata = 0
	v.CodeShare = 0
	v.DirectTransferType = 0
	v.Passenger = nil
	v.CreatePnrLimit = 0
	v.FarePrice = 0
	v.PriceType = 0
	v.OutMainFlight = false
	v.IsOwEoe = false
	poolPolicyDto.Put(v)
}
