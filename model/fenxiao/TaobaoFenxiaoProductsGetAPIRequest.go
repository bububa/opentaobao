package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductsGetAPIRequest
查询产品列表 API请求
taobao.fenxiao.products.get

查询供应商的产品数据。<br/><br/>    * 入参传入pids将优先查询，即只按这个条件查询。<br/>    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)<br/>    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。<br/>    * 入参fields传入images将查询多图数据，不传只返回主图数据。<br/>    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）<br/>    * 查询结果按照产品发布时间倒序，即时间近的数据在前。 */
type TaobaoFenxiaoProductsGetAPIRequest struct {
	model.Params
	// 商家编码
	_outerId string
	// 产品线ID
	_productcatId int64
	// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
	_pids string
	// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
	_fields string
	// 开始修改时间
	_startModified string
	// 结束修改时间
	_endModified string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// sku商家编码
	_skuNumber string
	// 查询产品列表时，查询入参“是否需要授权”<br/>yes:需要授权 <br/>no:不需要授权
	_isAuthz string
	// 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
	_itemIds string
}

// New
