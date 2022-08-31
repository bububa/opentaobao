package btrip

// OpenTrainOrderRs 结构体
type OpenTrainOrderRs struct {
	// 价目信息
	PriceInfoList []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>open_price_info,omitempty"`
	// 出行人列表
	UserAffiliateList []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list>open_user_affiliate_do,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 企业Id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 企业名称
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 联系人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 出发站
	DepStation string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// 到达站
	ArrStation string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 车次
	TrainNumber string `json:"train_number,omitempty" xml:"train_number,omitempty"`
	// 车次类型
	TrainType string `json:"train_type,omitempty" xml:"train_type,omitempty"`
	// 座位类型
	SeatType string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// 运行时长
	RunTime string `json:"run_time,omitempty" xml:"run_time,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 12306票号
	TicketNo12306 string `json:"ticket_no12306,omitempty" xml:"ticket_no12306,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 乘客姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 第三方行程id
	ThirdpartItineraryId string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// 订单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 票的数量
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 订单状态：0待支付,1出票中,2已关闭(未支付，已经关闭),3,改签成功,4退票成功,5出票完成,6退票申请中,7改签申请中,8已出票,已经邮寄了,9出票失败,10改签失败,11退票失败
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 成本中心对象
	CostCenter *OpenCostCenterDo `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 发票信息对象
	Invoice *OpenInvoiceDo `json:"invoice,omitempty" xml:"invoice,omitempty"`
}
