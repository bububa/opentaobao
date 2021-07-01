package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoOrderRemarkUpdateAPIRequest
修改采购单备注 API请求
taobao.fenxiao.order.remark.update

供应商修改采购单备注 */
type TaobaoFenxiaoOrderRemarkUpdateAPIRequest struct {
	model.Params
	// 采购单编号
	_purchaseOrderId int64
	// 备注内容(供应商操作)
	_supplierMemo string
	// 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
	_supplierMemoFlag int64
}

// New
