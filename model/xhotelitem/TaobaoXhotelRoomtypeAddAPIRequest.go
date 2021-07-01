package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomtypeAddAPIRequest
房型新增接口（ID重复变更新） API请求
taobao.xhotel.roomtype.add

房型添加或更新 */
type TaobaoXhotelRoomtypeAddAPIRequest struct {
	model.Params
	// 卖家房型ID，不能重复建议格式是:酒店code_房型code
	_outerId string
	// （已废弃）请使用outHid
	_hid int64
	// 房型名称。不能超过30字
	_name string
	// 最大入住人数，默认2（1-99）
	_maxOccupancy int64
	// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
	_area string
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
	_bedType string
	// 床宽。按自己定义存储，比如：2.1米
	_bedSize string
	// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
	_internet string
	// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
	_service string
	// 不要使用
	_extend string
	// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
	_windowType int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
	_srid int64
	// （必传）商家酒店ID，指明该房型属于哪家酒店
	_outHid string
	// 系统商，无申请不可使用
	_vendor string
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics string
	// 卖家房型英文名称
	_nameE string
	// 操作人信息
	_operator string
	// 属性值为1: 含义是非直连房型
	_connectionType int64
	// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
	_bedInfo string
	// 酒店房型设施
	_standardRoomFacilities string
}

// New
