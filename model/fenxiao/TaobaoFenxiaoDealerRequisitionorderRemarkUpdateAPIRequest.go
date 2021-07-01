package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest
修改经销采购单备注 API请求
taobao.fenxiao.dealer.requisitionorder.remark.update

供应商修改经销采购单备注 */
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest struct {
	model.Params
	// 经销采购单ID
	_dealerOrderId int64
	// 备注留言，可为空
	_supplierMemo string
	// 旗子的标记，必选。<br/>1-5之间的数字。<br/>非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
	_supplierMemoFlag int64
}

// New
