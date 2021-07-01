package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest
创建经销采购申请 API请求
taobao.fenxiao.dealer.requisitionorder.create

创建经销采购申请 */
type TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest struct {
	model.Params
	// 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货
	_logisticsType string
	// 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）
	_orderDetail []string
	// 收货人所在省份
	_province string
	// 收货人所在市
	_city string
	// 收货人所在区
	_district string
	// 收货人所在街道地址
	_address string
	// 收货人所在地区邮政编码
	_postCode string
	// 买家联系电话（此字段和mobile字段至少填写一个）
	_phone string
	// 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）
	_mobile string
	// 买家姓名（自提方式填写提货人姓名）
	_buyerName string
	// 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）
	_idCardNumber string
}

// New
