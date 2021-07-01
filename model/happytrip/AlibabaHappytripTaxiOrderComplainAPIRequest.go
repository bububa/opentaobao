package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderComplainAPIRequest
用户投诉 API请求
alibaba.happytrip.taxi.order.complain

一个订单只能投诉一次，不可重复投诉

投诉选项
0	其他原因；
1	司机原因导致行程被取消；
2	服务态度恶劣；
3	未坐车产生费用；
4	额外收取不合理费用；
5	路不熟多产生费用；
6	提前计费；
7	未及时结束计费；
8	司机绕路；
9	司机迟到；
10	司机爽约或拒载；
11	骚扰乘客；
12	危险驾驶；
13	不是订单显示车辆或司机； */
type AlibabaHappytripTaxiOrderComplainAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 投诉选项 0	其他原因； 1	司机原因导致行程被取消； 2	服务态度恶劣；	 3	未坐车产生费用；	 4	额外收取不合理费用； 5	路不熟多产生费用；	 6	提前计费； 7	未及时结束计费；	 8	司机绕路；	 9	司机迟到；	 10	司机爽约或拒载；	 11	骚扰乘客；	 12	危险驾驶；	 13	不是订单显示车辆或司机；
	_type int64
	// 投诉选项外的其他投诉内容,不能多于40个汉字
	_content string
	// 该用户的真实号码
	_mobile string
	// 投诉时间（毫秒）
	_time int64
}

// New
