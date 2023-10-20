package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductsgetAPIResponse 查询产品列表 API返回值
// taobao.fenxiao.products.get
//
// 查询供应商的产品数据。&lt;br/&gt;&lt;br/&gt;    * 入参传入pids将优先查询，即只按这个条件查询。&lt;br/&gt;    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)&lt;br/&gt;    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。&lt;br/&gt;    * 入参fields传入images将查询多图数据，不传只返回主图数据。&lt;br/&gt;    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）&lt;br/&gt;    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
type TaobaofenxiaoproductsgetAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoproductsgetAPIResponseModel
}

// TaobaofenxiaoproductsgetAPIResponseModel is 查询产品列表 成功返回结果
type TaobaofenxiaoproductsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
	Products []FenxiaoProduct `json:"products,omitempty" xml:"products>fenxiao_product,omitempty"`
	// 查询结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
