package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcEserviceAppointListAPIRequest
交互卡片预约信息读取接口 API请求
taobao.oc.eservice.appoint.list

允许外部的isv通过该接口读取门店预约信息 */
type TaobaoOcEserviceAppointListAPIRequest struct {
	model.Params
	// 预约信息唯一编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_code int64
	// 门店编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_mallCode string
	// 查询预约的起始时间，格式yyyyMMddHHmmss，默认为当前时间
	_startAppointTime string
	// 买家客户nick(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_customerNick string
	// 买家客户电话号码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_customerPhone string
	// 买家客户装修房屋所在的市(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_houseAddressCity string
	// 卖家主账号id
	_sellerId int64
	// 返回结果按预约时间排序，指示升序还是降息，取值asc和desc
	_sortOrder string
}

// New
